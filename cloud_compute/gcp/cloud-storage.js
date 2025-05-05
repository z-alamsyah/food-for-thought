import { Storage } from "@google-cloud/storage";

export default class CloudStorage {
  // The ID of your GCS bucket
  bucketName = "my-bucket-1357";
  //initialize storage
  storage = Storage;

  constructor(){
    this.storage = new Storage({
      keyFilename: "key.json",
    });
  }

   async uploadFile(file, fileTarget, folderTarget) {
      // The path to your file to upload
      const filePath = file;
  
      // The new ID for your GCS file
      const date = new Date();
      const destFileName = `${folderTarget}/${fileTarget}-` + date.toISOString();
  
      //upload doc
      const options = {
        destination: destFileName,
        preconditionOpts: {ifGenerationMatch: 0},
      };
  
      const resp = await this.storage.bucket(this.bucketName).upload(filePath, options);
      console.log(`${filePath} uploaded to ${this.bucketName}`);
      
      return destFileName;
    }

    async getSignedUrl(filePath, expiresInSeconds = 3600) {
      const options = {
        version: 'v4',
        action: 'read',
        expires: Date.now() + expiresInSeconds * 1000, // Default: 1 hour
      };
  
      const [url] = await this.storage
        .bucket(this.bucketName)
        .file(filePath)
        .getSignedUrl(options);
  
      return url;
    }
}

