import axios, { AxiosResponse } from "axios";
import { notification, message } from "antd";

export async function axiosPost(
  url: string,
  data: any,
  returnData: boolean = false,
  // config?: object,
  msg?: string
): Promise<any> {
  if (!msg) {
    msg = "posting request"
  }
  console.debug(data)
  const loadDone = message.loading(msg)
  let config = {
    validateStatus: validateStatus
  }
  const res = await axios.post(`${url}`, data, config)
    .then((res: AxiosResponse) => {
      loadDone()
      if (res.status >= 400) {
        message.error(msg + " failed: " + res.data)
        return false
      }
      if (res.status !== 200) {
        message.error(msg + ` unexpected code :${res.status} data:` + res.data)
        return false
      }
      message.success(msg + " success")
      return returnData ? res.data : true
    })
    .catch((reason: string) => {
      notification.error({
        message: msg + " failed",
        description:  reason,
      })
      return false
    })
  return res;
}

function validateStatus () {
  return true; // Reject only if the status code is greater than or equal to 500
}