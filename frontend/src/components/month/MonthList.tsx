import React, { useEffect, useState } from 'react';
import { worklog } from '../../../wailsjs/go/models';
import moment from 'moment';
type Props = {
  dates: worklog.WorkDay[];
};

const MonthList = ({ dates }: Props) => {
  const [index, setIndex] = useState<number>(0);
  const [keyPressed, setKeyPressed] = useState(false);

  useEffect(() => {
    dates.forEach((date) => {
      const momentDate = moment(date.date);
      const today = momentDate.isSame(new Date(), 'day');
      if (today) {
        setIndex(dates.indexOf(date));
      }
    });
  }, [dates]);

  //   React.useEffect(() => {
  //     const downHandler = (e: any) => {
  //       e.preventDefault();
  //       if (e.key === 'ArrowDown') {
  //         console.log(index);
  //         if (index) setIndex((ind) => ind + 1);
  //       }
  //     };

  //     const upHandler = (e: any) => {
  //       e.preventDefault();
  //       if (e.key === 'ArrowUp') {
  //         if (index && index >= 0) setIndex((ind) => ind - 1);
  //       }
  //     };

  //     window.addEventListener('keydown', downHandler);
  //     window.addEventListener('keyup', upHandler);

  //     return () => {
  //       window.removeEventListener('keydown', downHandler);
  //       window.removeEventListener('keyup', upHandler);
  //     };
  //   }, []);

  if (!index) {
    return null;
  }
  return (
    <div className='px-2 py-2 '>
      {dates.map((date, i) => {
        const momentDate = moment(date.date);

        return (
          <div className='py-0.5 cursor-pointer'>
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
