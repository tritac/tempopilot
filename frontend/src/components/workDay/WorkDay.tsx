import React, { useEffect, useState } from 'react';
import { useAppStore } from '../../AppContext';
import {
  DeleteWorkflow,
  GetWorkLog,
  PostWorkLog,
} from '../../../wailsjs/go/main/App';
import moment from 'moment';
import { worklog } from '../../../wailsjs/go/models';

type Props = {};

const initial = [
  { key: 'DEV', value: 0, color: '#1e40af' },
  { key: 'OTHER', value: 0, color: '#ae1eaf' },
  { key: 'SUPPORT', value: 0, color: '#af4e1e' },
  { key: 'ENABLERS', value: 0, color: '#1eafaa' },
  { key: 'COMSUP', value: 0, color: '#a5af1e' },
];
const WorkDay = (props: Props) => {
  const {
    appState: { selected, logsTypes, api },
  } = useAppStore();

  const [worklog, setWorklog] = useState<worklog.WorkLogResult[]>([]);
  const [disabled, setDisabled] = useState(false);
  const [total, setTotal] = useState(0);
  const [workLogType, setWorkLohType] =
    useState<{ key: string; value: number; color: string }[]>(initial);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setWorkLohType(initial);
    if (selected?.date) {
      console.log(moment(selected.date).unix());
      const unixTime = moment(selected.date).unix();
      setLoading(true);
      GetWorkLog(unixTime).then((res) => {
        setWorklog(res);
        console.log(res);
        setWorkLohType(initial);
        console.log(JSON.stringify(workLogType), 'tiy');
        const remote = workLogType.map((log) => {
          const remoteLog = res.find(
            (r) => r.attributes.values[0].value === log.key
          );
          if (remoteLog) {
            const hrs = remoteLog.timeSpentSeconds / (60 * 60);
            return { ...log, value: log.value + hrs };
          }
          // console.log(remoteLog);
          return { ...log, value: 0 };
        });

        setWorkLohType(remote);
        setLoading(false);
        // window.location.reload();
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

    if (!isNaN(value)) {
      setWorkLohType((e) =>
        e.map((s) => (s.key === key ? { ...s, key, value } : s))
      );
    }
  };
  const timeTime = moment(selected?.date).unix();
  const handlePostWork = () => {
    PostWorkLog(workLogType, timeTime)
      .then((res) => {
        console.log(res);
      })
      .then((e) => {
        console.log(e);
      });
  };

  const handleDeleteLog = (id: number) => {
    setLoading(true);
    DeleteWorkflow(id)
      .then((res) => {
        if (res) {
          const w = worklog.filter((e) => e.tempoWorklogId !== id);
          setWorklog(w);
          // console.log(w, 'W');
        }
        setLoading(false);
      })
      .catch((er) => {
        setLoading(false);
      });
  };

  return (
    <>
      <div className='flex flex-col '>
        <div
          className='border  flex items-center  px-2 h-10 flex-1'
          tabIndex={1}
        >
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
        <div className='w-full border   mt-4 px-2 relative  h-[calc(100vh-4.5rem)] '>
          <div className='grid grid-cols-2 text-xs py-2 justify-between'>
            <div className='flex  text-xs gap-x-3 '>
              <div className='bg-gray-700 w-16 px-2'>HOURS</div>
              <div>{total}</div>
            </div>
            <div className='flex'>
              <div className='bg-gray-700 w-16 px-2'>status:</div>
              <div>{loading ? 'Loading...' : ''}</div>
            </div>
            <div className='flex text-xs gap-x-3 col-span-3 items-center  '>
              <div className='bg-gray-700 w-16 px-2 mt-3'>APIKEY</div>
              <div>{api?.api_key}</div>
            </div>
          </div>
          <span className='text-xs bg-gray-950 -top-2.5 absolute left-2.5'>
            Work Log [ Ctrl + W ]
          </span>
          <div className='py-2 gap-x-1  flex  flex-col  justify-center items-center '>
            <div className='  gap-y-1 pt-9'>
              <div className='text-sm'>Hey {api?.user_name} , log work .</div>
              <div className='grid grid-cols-2 gap-x-2'>
                {workLogType.map((w) => {
                  return (
                    <div className='flex  my-3' key={w.key}>
                      <div
                        className='px-2 w-24  text-sm flex items-center'
                        style={{ backgroundColor: w.color }}
                      >
                        {w.key}
                      </div>
                      {}
                      <div>
                        <input
                          type='number'
                          max={8}
                          min={0}
                          value={
                            workLogType.find((s) => s.key === w.key)?.value
                          }
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
                onClick={() => handlePostWork()}
                disabled={disabled}
              >
                SUBMIT
              </button>
            </div>
          </div>
          <div>
            <div className='grid grid-cols-4 w-full '>
              {worklog.map((re) => {
                return (
                  <div
                    className='text-xs  my-3 flex w-full justify-between '
                    key={re.tempoWorklogId}
                  >
                    <div className='bg-gray-700  w-1/2'>
                      {re.attributes.values[0].value}
                    </div>
                    <div className='pl-2 '>{re.timeSpentSeconds / 3600}</div>
                    <div
                      onClick={() => handleDeleteLog(re.tempoWorklogId)}
                      className='bg-red-400 px-2 mx-2 self-end cursor-pointer'
                    >
                      x
                    </div>
                  </div>
                );
              })}
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default WorkDay;
