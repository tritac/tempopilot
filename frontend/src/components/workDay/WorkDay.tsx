import React, { useEffect, useState } from 'react';
import { useAppStore } from '../../AppContext';
import { GetWorkLog } from '../../../wailsjs/go/main/App';
import moment from 'moment';
import { worklog } from '../../../wailsjs/go/models';

type Props = {};

const WorkDay = (props: Props) => {
  const {
    appState: { selected, logsTypes, api },
  } = useAppStore();

  const [worklog, setWorklog] = useState<worklog.WorkLogResult[]>([]);
  const [disabled, setDisabled] = useState(false);
  const [total, setTotal] = useState(0);
  const [workLogType, setWorkLohType] = useState<
    { key: string; value: number; color: string }[]
  >([
    { key: 'DEV', value: 0, color: '#1e40af' },
    { key: 'OTHER', value: 0, color: '#ae1eaf' },
    { key: 'SUPPORT', value: 0, color: '#af4e1e' },
    { key: 'ENABLERS', value: 0, color: '#1eafaa' },
    { key: 'COMSUP', value: 0, color: '#a5af1e' },
  ]);

  useEffect(() => {
    if (selected?.date) {
      console.log(moment(selected.date).unix());
      const unixTime = moment(selected.date).unix();
      console.log(unixTime, 'nf');
      GetWorkLog(unixTime).then((res) => {
        res.map((log) => {
          const hrs = log.timeSpentSeconds / (60 * 60);
          const key = log.attributes.values[0].value;
          console.log(key);
          console.log(hrs);
          const remote = workLogType.map((el) =>
            el.key === key ? { ...el, value: el.value + hrs } : el
          );
          console.log(remote, 're');
          setWorkLohType(remote);
        });
        setWorklog(res);
      });
    }
  }, [selected]);

  useEffect(() => {
    const total = workLogType.reduce((a, b) => a + b.value, 0);
    setTotal(total);
    if (total > 8 || total === 0) {
      setDisabled(true);
    } else {
      setDisabled(false);
    }
  }, [workLogType]);

  const handleChange = (key: string, value: number) => {
    // const total = workLogType.reduce((a, b) => a + b.value, 0);

    console.log(key, value);
    if (!isNaN(value)) {
      setWorkLohType((e) =>
        e.map((s) => (s.key === key ? { ...s, key, value } : s))
      );
    }
  };

  return (
    <>
      <div className='border  flex items-center  px-2 h-10'>
        {workLogType.map((wl) => (
          <div
            className='h-full '
            style={{
              backgroundColor: wl.color,
              width: `${(wl.value / 8) * 100}%`,
            }}
          ></div>
        ))}

        <span className='absolute top-0 bg-gray-950  text-xs left-5'>
          Day Log
        </span>
      </div>

      <div className='w-full border   mt-4 px-2 relative h-[calc(100%-4.5rem)] '>
        <span className='text-xs bg-gray-950 -top-2.5 absolute left-2.5'>
          Work Log [ Ctrl + W ]
        </span>
        <div className='py-2 gap-x-1  flex  flex-col  justify-center items-center '>
          <div className='  gap-y-1 py-9'>
            <div className='text-sm'>
              Hey {api?.user_name} today your total working hours are {total}
            </div>
            <div className='grid grid-cols-2 gap-x-2'>
              {workLogType.map((w) => {
                return (
                  <div className='flex  my-3'>
                    <div className='px-2 w-24 bg-slate-700 text-sm flex items-center'>
                      {w.key}
                    </div>
                    {}
                    <div>
                      <input
                        type='number'
                        max={8}
                        min={0}
                        value={workLogType.find((s) => s.key === w.key)?.value}
                        className=' bg-slate-500 outline-none border-[1px] w-28 px-1'
                        // disabled={}
                        onChange={(e) => {
                          handleChange(w.key, parseFloat(e.target.value));
                        }}
                      />
                    </div>
                  </div>
                );
              })}
            </div>
            <button
              className={`w-full  text-center outline-[0.5px] ${
                disabled ? 'bg-gray-400' : 'bg-green-500'
              }  `}
              disabled={disabled}
            >
              SUBMIT
            </button>
          </div>
        </div>
      </div>
    </>
  );
};

export default WorkDay;
