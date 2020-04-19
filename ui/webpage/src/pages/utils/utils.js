import axios from 'axios';

export const apiGet = async (path) => {
  let ret = await axios({
    method: 'GET',
    url: global.server + path,
    timeout: 20000,
  });
  return ret
}

export const apiPost = async (path, data) => {
  let ret = await axios({
    method: 'POST',
    url: global.server + path,
    timeout: 20000,
    data
  });
  return ret
}

export const apiAsyncGet = (path, callback) => {
  axios({
    method: 'GET',
    url: global.server + path,
    timeout: 20000,
  }).then(callback).catch(console.log);
}

export const apiAsyncPost = (path, data, callback) => {
  axios({
    method: 'POST',
    url: global.server + path,
    timeout: 20000,
    data
  }).then(callback).catch(console.log);
}