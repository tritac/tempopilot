import { useEffect, useState } from 'react';
import DayLog from './components/daylog/DayLog';
import { Environment, EventsEmit } from '../wailsjs/runtime/runtime';
import Intro from './components/dialog/Intro';
import { GetUserConfig, GetMonthList } from '../wailsjs/go/main/App';
import { appstore } from '../wailsjs/go/models';
import { useAppStore } from './AppContext';
import Month from './components/month/Month';
import WorkDay from './components/workDay/WorkDay';

function App() {
  const {
    appState: { api },
  } = useAppStore();

  console.log(api);
  if (!api?.isValidApi) {
    return <Intro />;
  }

  return (
    <div id='App '>
      <div className='flex '>
        <div className='w-1/3 bg-black text-white h-screen border-r-[1px] '>
          <Month />
        </div>
        <div className='w-full bg-gray-950 text-white pt-2 px-2 relative'>
          <div className='h-10'>
            <DayLog />
          </div>
          <div className='w-full border  h-[calc(100%-4rem)] mt-4 px-2 relative'>
            <span className='text-xs bg-gray-950 -top-2.5 absolute left-2.5'>
              Work Log [ Ctrl + W ]
            </span>
            <div>
              <WorkDay />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
