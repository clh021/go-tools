
export function formatBytes(bytes, decimals = 2, separate = false) {
  if (bytes === 0) return "0 B";

  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

  const i = Math.floor(Math.log(bytes) / Math.log(k));
  const val = parseFloat((bytes / Math.pow(k, i)).toFixed(dm));
  const unit = sizes[i];
  let res = null;
  if (separate) {
    res = { val: val, unit: unit };
  } else {
    res = val + " " + unit;
  }
  return res;
}
