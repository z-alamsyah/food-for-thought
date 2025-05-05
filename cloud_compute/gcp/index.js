import CloudStorage from "./cloud-storage.js";

const storage = new CloudStorage();


// upload images
const objFile = await storage.uploadFile('dart.jpeg', 'my-image', 'images');
const signURL = await storage.getSignedUrl(objFile, 120);


//For Access signURL, you can use <img src="" />
//<img src="https://storage.googleapis.com/my-bucket-1357/images/my-image-...signed-part..." />
console.log('Access URL:', signURL);
