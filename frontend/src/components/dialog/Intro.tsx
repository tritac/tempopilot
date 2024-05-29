import React, { useState } from 'react';
import { CreateUserConfig, VerifyApiKey } from '../../../wailsjs/go/main/App';
import { useAppStore } from '../../AppContext';
type Props = {};

function Intro({}: Props) {
  const {
    appState: { api },
  } = useAppStore();
  const [userName, setUserName] = useState('');
  const [tempoApiKey, setApiKey] = useState('');
  const [tempoUserId, setTempoId] = useState('');

  const submitHandler = () => {
    // CreateUserConfig(userName, tempoApiKey, tempoUserId);
    VerifyApiKey(tempoApiKey)
      .then((res) => {
        if (res) {
          CreateUserConfig(userName, tempoApiKey, tempoUserId)
            .then((res) => {
              location.reload();
            })
            .then((err) => {});
        }
      })
      .then((err: any) => {
        console.log(err, 'Error');
      });
  };

  return (
    <div className='flex h-screen w-full bg-black  text-white '>
      <div className='border w-full mx-2 my-2 flex justify-center items-center '>
        <div className='flex flex-col gap-3'>
          {JSON.stringify(api)}
          <div className='relative  flex flex-col'>
            <span className='text-xs absolute -top-2 bg-black left-3 px-1'>
              Name
            </span>
            <input
              type='text'
              value={userName}
              className=' border outline-none bg-black pl-2 py-1'
              onChange={(e) => setUserName(e.target.value)}
            />
          </div>
          <div className='relative  flex flex-col'>
            <span className='text-xs absolute -top-2 bg-black left-3 px-1'>
              Tempo Api Key
            </span>
            <input
              type='text'
              value={tempoApiKey}
              className=' border outline-none bg-black pl-2 py-1'
              onChange={(e) => setApiKey(e.target.value)}
            />
          </div>
          <div className='relative  flex flex-col'>
            <span className='text-xs absolute -top-2 bg-black left-3 px-1'>
              Tempo User Id
            </span>
            <input
              type='text'
              value={tempoUserId}
              className=' border outline-none bg-black pl-2 py-1'
              onChange={(e) => setTempoId(e.target.value)}
            />
          </div>

          <button className='border bg-green-500' onClick={submitHandler}>
            SUBMIT
          </button>
          <button className='border bg-red-500'>CANCEL</button>
        </div>
      </div>
    </div>
  );
}

export default Intro;
