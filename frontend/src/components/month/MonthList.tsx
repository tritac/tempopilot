import React, { useEffect, useState } from 'react';
import { worklog } from '../../../wailsjs/go/models';
import moment from 'moment';
import { useAppStore } from '../../AppContext';
type Props = {
  dates: worklog.WorkDay[];
};

const MonthList = ({ dates }: Props) => {
  const [index, setIndex] = useState<number>(0);
  const { dispatch } = useAppStore();

  useEffect(() => {
    dates.forEach((date) => {
      const momentDate = moment(date.date);
      const today = momentDate.isSame(new Date(), 'day');
      if (today) {
        setIndex(dates.indexOf(date));
      }
    });
  }, [dates]);

  const handleDateSelect = (data: worklog.WorkDay, index: number) => {
    setIndex(index);
    dispatch({ type: 'SET_SELECTED_DATE', payload: data });
  };

  return (
    <div className='px-2 py-2 '>
      {dates.map((date, i) => {
        const momentDate = moment(date.date);

        return (
          <div
            className='py-0.5 cursor-pointer'
            key={date.date}
            onClick={() => handleDateSelect(date, i)}
          >
            <div
              className={`text-sm flex gap-x-2 ${
                i == index ? 'bg-teal-800' : ''
              }`}
            >
              <div
                className={` ${
                  date.day === 'Saturday' || date.day === 'Sunday'
                    ? 'bg-red-700'
                    : 'bg-gray-700'
                }  text-xs flex items-center justify-center w-4`}
                tabIndex={i}
              >
                {date.day.substring(0, 1)}
              </div>
              <div>{momentDate.format('MMMM Do YYYY')}</div>
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default MonthList;
