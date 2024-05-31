import React, { useEffect, useState } from 'react';
import MonthList from './MonthList';
import { useAppStore } from '../../AppContext';
import { GetMonthList } from '../../../wailsjs/go/main/App';
import { worklog } from '../../../wailsjs/go/models';
import moment, { now } from 'moment';

type Props = {};

const Month = (props: Props) => {
  const {
    appState: { api },
  } = useAppStore();

  const [dateList, setDateList] = useState<worklog.WorkDay[]>([]);

  useEffect(() => {
    if (api?.isValidApi) {
      const time = moment(now()).month();
      console.log({ time });
      GetMonthList(2024, time + 1).then((res) => {
        setDateList(res);
        // console.log(res);
      });
    }
  }, [api]);

  return (
    <div className='relative'>
      <div className=' relative h-[calc(100vh-10vh)] overflow-auto'>
        <div>
          <MonthList dates={dateList} />
          {/* <div>as</div> */}
        </div>
      </div>
      <div className='absolute px-2 flex justify-between w-full pt-2 text-sm'>
        <div className='uppercase bg-teal-700 px-2'>next</div>
        <div className='uppercase bg-teal-700 px-2'>Prev</div>
      </div>
    </div>
  );
};

export default Month;
