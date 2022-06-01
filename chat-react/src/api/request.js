import axios from "axios";
import qs from "qs";

// 创建 axios 实例
const request = axios.create({
  // API 请求的默认前缀
  baseURL: "/api",
});

// request interceptor
request.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    // do something with request error
    console.log(error); // for debug
    return Promise.reject(error);
  }
);

// response interceptor
request.interceptors.response.use(
  (response) => {
    // return response.data //可以不做任何处理，在这里了直接返回请求会的结果
    const res = response.data || {};
    return res;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default {
  /**
   *
   * @param {String} url
   * @param {Object} data
   * @param {String} isForm
   * @param {Object} config
   *
   * 这里是通用请求
   * 如果需要特殊的配置在config中配置
   *
   */
  post(url, data, isForm, config = {}) {
    if (isForm) {
      const transformRequest = (data) =>
        qs.stringify(data, {
          encode: false,
          allowDots: true,
          arrayFormat: "indices",
        });
      config.transformRequest = [transformRequest];
    }
    // 文件上传
    if (isForm === "FILE") {
      const formHeaders = {
        "Content-Type": "multipart/form-data",
      };
    }
    return request.post(url, data, config);
  },
  /**
   *
   * @param {String} url
   * @param {Object} params
   * @param config
   */
  get(url, params, config) {
    return request.get(url, { params, ...config });
  },
};
