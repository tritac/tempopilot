import React, { useState } from 'react';
import { CreateUserConfig, VerifyApiKey } from '../../../wailsjs/go/main/App';
import { useAppStore } from '../../AppContext';
type Props = {
  setShowSetting: React.Dispatch<React.SetStateAction<boolean>>;
};

function Intro({ setShowSetting }: Props) {
  const {
    appState: { api },
  } = useAppStore();
  const [userName, setUserName] = useState(api?.user_name || '');
  const [tempoApiKey, setApiKey] = useState(api?.api_key || '');
  const [tempoUserId, setTempoId] = useState(api?.user_id || '');

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
          {/* {JSON.stringify(api)} */}
          <div className='text-sm space-y-2'>
            <div className='flex'>
              <div className='bg-gray-800 px-2 w-20 text-xs'>NAME</div>
              <div className='px-2'>{api?.user_name}</div>
            </div>
            <div className='flex'>
              <div className='bg-gray-800 px-2 w-20 text-xs'>API KEY</div>
              <div className='px-2'>{api?.api_key}</div>
            </div>
            <div className='flex'>
              <div className='bg-gray-800 px-2 w-20 text-xs'>WORKER ID</div>
              <div className='px-2'>{api?.user_id}</div>
            </div>
            <div className='flex'>
              <div className='bg-gray-800 px-2 w-20 text-xs'>API VALID</div>
              <div className='px-2 flex items-center'>
                {api?.isValidApi ? (
                  <div className='h-2 w-2 rounded-sm bg-green-500'></div>
                ) : (
                  <div className='h-2 w-2 rounded-sm bg-red-500'></div>
                )}
              </div>
            </div>
          </div>
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
          <button
            onClick={() => setShowSetting(false)}
            className='border bg-red-500'
          >
            CANCEL
          </button>
        </div>
      </div>
    </div>
  );
}

export default Intro;
