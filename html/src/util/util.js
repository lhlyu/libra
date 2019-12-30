const _defaultUrl = "https://cdn.jsdelivr.net/gh/lhlyu/pb@master/2019/6.jpg";

function loadImageAsync(url) {
  return new Promise((resolve)=> {
      let image = new Image();
      image.onload = function() {
          resolve(image);
      };
      image.onerror = function() {
          image.src = _defaultUrl
          resolve(image);
      }
      image.src = url;
  })
}


export default {
    loadImageAsync
}
