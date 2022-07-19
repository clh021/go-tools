package uploadAdvanced

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func Main() {
	port := 8080
	router := gin.Default()

	router.Use(Cors)
	router.GET("/checkChunk", func(c *gin.Context) {
		hash := c.Query("hash")
		hashPath := fmt.Sprintf("./uploadFile/%s", hash)
		chunkList := []string{}
		isExistPath, err := PathExists(hashPath)
		if err != nil {
			fmt.Println("获取hash路径错误", err)
		}
		if isExistPath {
			files, err := ioutil.ReadDir(hashPath)
			state := 0
			if err != nil {
				fmt.Println("文件读取错误", err)
			}
			for _, f := range files {
				fileName := f.Name()
				chunkList = append(chunkList, fileName)
				fileBaseName := strings.Split(fileName, ".")[0]
				if fileBaseName == hash {
					state = 1
				}
			}

			c.JSON(200, gin.H{
				"state":     state,
				"chunkList": chunkList,
			})
		} else {
			os.MkdirAll(hashPath, os.ModePerm)
			c.JSON(200, gin.H{
				"state":     0,
				"chunkList": chunkList,
			})
		}
		if len(chunkList) > 0 {
			fmt.Printf("http://%s/uploadFile/%s/%s\n", c.Request.Host, hash, chunkList[len(chunkList)-1])
		}
	})

	router.POST("/uploadChunk", func(c *gin.Context) {
		fileHash := c.PostForm("hash")
		file, err := c.FormFile("file")
		hashPath := fmt.Sprintf("./uploadFile/%s", fileHash)
		if err != nil {
			fmt.Println("获取上传文件失败", err)
		}

		isExistPath, err := PathExists(hashPath)
		if err != nil {
			fmt.Println("获取hash路径错误", err)
		}

		if !isExistPath {
			os.MkdirAll(hashPath, os.ModePerm)
		}

		err = c.SaveUploadedFile(file, fmt.Sprintf("./uploadFile/%s/%s", fileHash, file.Filename))
		if err != nil {
			c.String(400, "0")
			fmt.Println(err)
		} else {
			chunkList := []string{}
			files, err := ioutil.ReadDir(hashPath)
			if err != nil {
				fmt.Println("文件读取错误", err)
			}
			for _, f := range files {
				fileName := f.Name()

				if f.Name() == ".DS_Store" {
					continue
				}
				chunkList = append(chunkList, fileName)
			}

			c.JSON(200, gin.H{
				"chunkList": chunkList,
				"fileUrl":   fmt.Sprintf("http://%s/uploadFile/%s/%s", c.Request.Host, fileHash, file.Filename),
			})
		}
	})

	router.GET("/megerChunk", func(c *gin.Context) {
		hash := c.Query("hash")
		fileName := c.Query("fileName")
		hashPath := fmt.Sprintf("./uploadFile/%s", hash)

		isExistPath, err := PathExists(hashPath)
		if err != nil {
			fmt.Println("获取hash路径错误", err)
		}

		if !isExistPath {
			c.JSON(400, gin.H{
				"message": "文件夹不存在",
			})
			return
		}
		isExistFile, err := PathExists(hashPath + "/" + fileName)
		if err != nil {
			fmt.Println("获取hash路径文件错误", err)
		}
		fmt.Println("文件是否存在", isExistFile)
		if isExistFile {
			c.JSON(200, gin.H{
				"fileUrl": fmt.Sprintf("http://%s/uploadFile/%s/%s", c.Request.Host, hash, fileName),
			})
			return
		}

		files, err := ioutil.ReadDir(hashPath)
		if err != nil {
			fmt.Println("合并文件读取失败", err)
		}
		complateFile, err := os.Create(hashPath + "/" + fileName)
		defer complateFile.Close()
		for _, f := range files {
			//.DS_Store
			//file, err := os.Open(hashPath + "/" + f.Name())
			//if err != nil {
			//	fmt.Println("文件打开错误", err)
			//}

			if f.Name() == ".DS_Store" {
				continue
			}

			fileBuffer, err := ioutil.ReadFile(hashPath + "/" + f.Name())
			if err != nil {
				fmt.Println("文件打开错误", err)
			}
			complateFile.Write(fileBuffer)
		}

		c.JSON(200, gin.H{
			"fileUrl": fmt.Sprintf("http://%s/uploadFile/%s/%s", c.Request.Host, hash, fileName),
		})

	})

	router.GET("/", func(c *gin.Context) {
		fmt.Println("c.Request.Host", c.Request.Host, c.Request.RequestURI, c.Request.RemoteAddr, c.Request.URL)
		html := `<script src="https://cdn.bootcdn.net/ajax/libs/spark-md5/3.0.0/spark-md5.min.js"></script>
<br>	TODO: 数据上传触发事件
<br>	<input type="file" id="file"><button onclick="upload()">upload</button><script>
        const sliceSingleSize = 1024 * 1024 * 2;
        document.querySelector('input').onchange = function(e) {
			console.log('trigger upload not on this change event. use button.')
		}
		function upload() {
			const file = document.getElementById("file").files[0];
            const sliceBuffer = []
            let sliceSize = file.size
            while(sliceSize > sliceSingleSize) {
                const blobPart = file.slice(sliceBuffer.length * sliceSingleSize, (sliceBuffer.length + 1) * sliceSingleSize)
                sliceBuffer.push(
                    blobPart
                )
                sliceSize -= sliceSingleSize
            }

            if(sliceSize > 0) {
                sliceBuffer.push(
                    file.slice(sliceBuffer.length * sliceSingleSize, file.size)
                )
            }
            
            const fileReader = new FileReader()
            fileReader.onload = function(res){
                const result = fileReader.result
                const fileHash = SparkMD5.hashBinary(result)

                checkFileChunkState(fileHash)
                .then(res => {
                    let { chunkList, state } = res
                    if(state === 1) {
                        alert("已经上传完成")
                        return 
                    }

                    chunkList = chunkList.map(e => parseInt(e))

                    const chunkRequests = []
                    sliceBuffer.forEach((buffer, i) => {
                        if(!chunkList.includes(i)) {
                            const blob = new File([buffer], i)
							let percent = Math.round((i+1)/sliceBuffer.length * 10000) / 100 + "%";
							console.log("before uploadFileChunk:", i, sliceBuffer.length, percent)
                            chunkRequests.push(
                                uploadFileChunk(fileHash, blob)
                            )
                        }
                    })
                    return Promise.all(chunkRequests)
                })
                .then(res => {
                    return new Promise(resolve => {
                        res.forEach((e, i) => {
							let percent = Math.round((i+1)/res.length * 10000) / 100 + "%";
							console.log("before megerChunkFile:", i, res.length, percent)
                            e.json().then(({chunkList}) => {
                                if(chunkList.length === sliceBuffer.length) {
                                    megerChunkFile(fileHash, file.name).then(res => {
                                        resolve(res)
                                    })
                                }
                            })
                        })
                    })
                }).then(res => {
                    console.log(res)
                })
            }
            fileReader.onerror = function(err) {
                console.log("报错了", err.target.error)
            }
            fileReader.readAsBinaryString(file)

        }

        function uploadFileChunk(hash, file) {
            let formData = new FormData
            formData.append('file', file)
            formData.append('hash', hash)
            return fetch("http://` + c.Request.Host + `/uploadChunk", {
                method: "POST",
                body: formData
            })
        }

        function checkFileChunkState(hash) {
            return new Promise(resolve => {
                fetch("http://` + c.Request.Host + `/checkChunk?hash=" + hash)
                .then(r => r.json())
                .then(response => {
                    resolve(response)
                })
            })
        }

        function megerChunkFile(hash, fileName) {
            return new Promise(resolve => {
                fetch('http://` + c.Request.Host + `/megerChunk?hash=' + hash + '&fileName=' + fileName)
                .then(r => r.json())
                .then(r => {
                    resolve(r)
                })
            })
        }

    </script>`
		c.Data(http.StatusOK, ContentTypeHTML, []byte(html))
	})

	router.StaticFS("/uploadFile", http.Dir("uploadFile"))
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("router error: ", err)
		os.Exit(1)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Cors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
