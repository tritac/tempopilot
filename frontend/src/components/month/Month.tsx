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
    dispatch,
  } = useAppStore();

  const [dateList, setDateList] = useState<worklog.WorkDay[]>([]);
  const [month, setMonth] = useState(moment(new Date()).month() + 1);
  useEffect(() => {
    if (api?.isValidApi) {
      GetMonthList(2024, month).then((res) => {
        setDateList(res);
        // console.log(res);
      });
    }
  }, [api, month]);

  const handlePreviousMonth = () => {
    setMonth((s) => s - 1);
  };
  const handleNextMonth = () => {
    setMonth((s) => s + 1);
  };

  return (
    <div className='relative'>
      <div className=' relative h-[calc(100vh-10vh)] overflow-auto'>
        <div>
          <MonthList dates={dateList} />
          {/* <div>as</div> */}
        </div>
      </div>
      <div className='absolute px-2 flex justify-between w-full pt-2 text-sm'>
        <div
          className='uppercase bg-teal-700 px-2 cursor-pointer'
          onClick={handlePreviousMonth}
        >
          Prev
        </div>
        <div
          className='uppercase bg-teal-700 px-2 cursor-pointer'
          onClick={handleNextMonth}
        >
          Next
        </div>
      </div>
    </div>
  );
};

export default Month;
