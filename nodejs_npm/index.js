import * as dotenv from 'dotenv';
import axios from 'axios';
dotenv.config();

//Contoh Penggunaan ENV
console.info("Ini isi ENV -->>", process.env.TOKEN);

//Contoh Penggunaan Axios
const response = await axios.get(process.env.BASE_URL + '/en/characters');
console.info(response.data);


