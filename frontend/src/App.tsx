import { useEffect, useState } from 'react';
import DayLog from './components/daylog/DayLog';
import { Environment, EventsEmit, EventsOn } from '../wailsjs/runtime/runtime';
import Intro from './components/dialog/Intro';
import { GetUserConfig, GetMonthList } from '../wailsjs/go/main/App';
import { appstore } from '../wailsjs/go/models';
import { useAppStore } from './AppContext';
import Month from './components/month/Month';
import WorkDay from './components/workDay/WorkDay';

function App() {
  const [showSetting, setShowSetting] = useState(false);
  const {
    appState: { api },
  } = useAppStore();
  const [apiValid, setApiValid] = useState(true);

  console.log(api);

  useEffect(() => {
    EventsOn('SETTING', () => {
      setShowSetting(!showSetting);
    });

    EventsOn('INVALIDATE', () => {
      setApiValid(false);
    });
  }, []);

  if ((!api?.isValidApi && apiValid) || showSetting) {
    return (
      <>
        {/* <div className='text-white'>{apiValid ? 'true' : 'false'}</div> */}
        <Intro setShowSetting={setShowSetting} />
      </>
    );
  }

  const handleReload = () => {
    EventsEmit('RELOAD');
  };

  return (
    <div id='App '>
      <div className='flex '>
        <div className='w-1/3 bg-black text-white h-screen border-r-[1px]  relative'>
          <Month />
        </div>
        <div className='w-full bg-gray-950 text-white pt-2 px-2 relative'>
          {/* <div className='h-10'>
            <DayLog />
          </div> */}
          {/* <div className='w-full border  h-[calc(100%-4rem)] mt-4 px-2 relative'> */}
          <div className='relative'>
            <div
              className='absolute z-30 cursor-pointer top-14 right-2 px-1 py-1 bg-blue-700  text-xs'
              onClick={handleReload}
            >
              Reload{' '}
            </div>
            <WorkDay />
          </div>
          {/* </div> */}
        </div>
      </div>
    </div>
  );
}

export default App;
