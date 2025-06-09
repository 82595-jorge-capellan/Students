import { check, sleep } from 'k6';
import grpc from 'k6/net/grpc';
import encoding from 'k6/encoding';
import http from 'k6/http';
import { Trend } from 'k6/metrics';

const client = new grpc.Client();
client.load(['.'], '../protobuf/School.proto');

function randomIntFromInterval(min, max) { // min and max included 
  return Math.floor(Math.random() * (max - min + 1) + min);
}



export const options = {
  insecureSkipTLSVerify: true,
  stages: [
    {
      duration: "60s",
      target: 100,
    },{
      duration: "120s",
      target: 100,
    }
  ]
}

export function setup() {
  const url = 'https://localhost:9200/_cache/clear';
  const user = 'admin';
  const pass = 'Opensearch1234*';
  const credentials = `${user}:${pass}`;
  const encodedCredentials = encoding.b64encode(credentials);

  const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Basic ${encodedCredentials}`,
  };

  const res = http.post(url, null, { headers, tags: { name: 'clear-cache' } });

  if (res.status !== 200) {
    throw new Error('Error limpiando cache: ' + res.status + ' ' + res.body);
  }

  console.log('Cache limpiado correctamente');
}


export default () => {
  client.connect('localhost:50051', {
    plaintext: true
  });

  let sid = randomIntFromInterval(1, 1000);
  const data = { id: sid, };

  const impl = __ENV.IMPL || 'MS';

  const response = client.invoke(`data.School/SearchStudentByID${impl}`, data);



  check(response, {
    'status is OK': (r) => r && r.status === grpc.StatusOK,
  });

  client.close();
  sleep(1);
};