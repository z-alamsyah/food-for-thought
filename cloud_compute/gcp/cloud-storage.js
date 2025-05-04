import { Storage } from "@google-cloud/storage";

export default class CloudStorage {
   async uploadFile() {
      const storage = new Storage({
        keyFilename: "key.json",
      });
  
      // The ID of your GCS bucket
      const bucketName = "my-bucket-1357";
  
      // The path to your file to upload
      const filePath = 'test-doc-csv.csv';
  
      // The new ID for your GCS file
      const date = new Date();
      const destFileName = 'docs/my-file-csv-' + date.toISOString();
  
      //upload doc
      const options = {
        destination: destFileName,
        preconditionOpts: {ifGenerationMatch: 0},
      };
  
      await storage.bucket(bucketName).upload(filePath, options);
      console.log(`${filePath} uploaded to ${bucketName}`);
    }
}

