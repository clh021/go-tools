import axios from "axios";
import SparkMd5 from "spark-md5";

async function getChunksFromFile(file, pieceSize, cb) {
  const chunkSize = pieceSize * 1024 * 1024; // 5MM
  const chunkCount = Math.ceil(file.size / chunkSize); // 总片数
  for (let index = 0; index < chunkCount; index++) {
    let start = index * chunkSize;
    let end = Math.min(file.size, start + chunkSize);
    let chunk = file.slice(start, end);
    const res = await cb({ chunk, index, total: chunkCount });
    if (!res) {
      break;
    }
  }
}

// 分片上传
/*
 * 分片上传函数 支持多个文件
 * @param options
 * options.file 表示源文件
 * options.pieceSize 表示需要分片的大小 默认是5m
 * options.chunkUrl 分片上传的后端地址
 * options.fileUrl 整个文件的上传地址
 * progress 进度回调
 * success 成功回调
 * error 失败回调
 *
 */
export const uploadByPieces = ({
  file,
  chunkUrl,
  fileUrl,
  pieceSize = 5,
  progress,
  success,
  error,
}) => {
  // 上传过程中用到的变量
  let progressNum = 1; // 进度
  let successAllCount = 0; // 上传成功的片数
  // let currentAllChunk = 0 // 当前上传的片数索引
  // 获取md5
  const readFileMD5 = async () => {
    progress(0);
    // 读取每个文件的md5
    const hash = new SparkMd5.ArrayBuffer();

    await getChunksFromFile(file, 5, async ({ chunk, index, total }) => {
      let chunkFR = new FileReader();
      return await new Promise((resolve) => {
        chunkFR.readAsArrayBuffer(chunk);
        chunkFR.onload = (e) => {
          hash.append(e.target.result);
          resolve(true);
        };
      });
    });

    await readChunkMD5({ md5: hash.end(), name: file.name, file });
  };
  // 针对每个文件进行chunk处理
  const readChunkMD5 = async (currentFile) => {
    let chunkCount;
    await getChunksFromFile(
      currentFile.file,
      pieceSize,
      async ({ chunk, index, total }) => {
        chunkCount = total;
        let chunkFR = new FileReader();
        let chunkMD5;
        const hash = new SparkMd5.ArrayBuffer();
        return await new Promise((resolve, reject) => {
          chunkFR.readAsArrayBuffer(chunk);
          chunkFR.addEventListener(
            "load",
            (e) => {
              hash.append(e.target.result);
              chunkMD5 = hash.end();
              uploadChunk(currentFile, {
                chunkMD5,
                chunk,
                currentChunk: index,
                chunkCount: total,
              })
                .then((resp) => {
                  resolve(resp.code != 1);
                })
                .catch((e) => {
                  reject(e);
                });
            },
            false
          );
        });
      }
    );

    uploadFile(currentFile, chunkCount);
  };

  const uploadChunk = (currentFile, chunkInfo) => {
    let fetchForm = new FormData();
    fetchForm.append("file_name", currentFile.name);
    fetchForm.append("md5", currentFile.md5);
    fetchForm.append("file", chunkInfo.chunk);
    fetchForm.append("chunks", chunkInfo.chunkCount);
    fetchForm.append("chunk_index", chunkInfo.currentChunk);
    fetchForm.append("chunk_md5", chunkInfo.chunkMD5);
    return axios({
      method: "post",
      url: chunkUrl,
      data: fetchForm,
    }).then((res) => {
      progressFun(chunkInfo.currentChunk, chunkInfo.chunkCount);
      if (successAllCount < chunkInfo.chunkCount - 1) {
        successAllCount++;
      } else {
      }
      return res;
    });
  };

  const progressFun = (current, total) => {
    progressNum = Math.ceil((current / total) * 100);
    progress(progressNum);
  };
  // 对分片已经处理完毕的文件进行上传
  const uploadFile = (currentFile, chunkCount) => {
    let makeFileForm = new FormData();
    makeFileForm.append("md5", currentFile.md5);
    makeFileForm.append("count", chunkCount);
    makeFileForm.append("file_name", currentFile.name);
    makeFileForm.append("size", currentFile.file.size);
    return axios({
      // 合并文件
      method: "post",
      url: fileUrl,
      data: makeFileForm,
    })
      .then((res) => {
        success && success(res);
        successAllCount++;
      })
      .catch((e) => {
        error && error(e);
      });
  };
  readFileMD5(); // 开始执行代码
};
