import React, { useEffect, useState } from 'react';
import { useAppStore } from '../../AppContext';
import { GetWorkLog } from '../../../wailsjs/go/main/App';
import moment from 'moment';
import { worklog } from '../../../wailsjs/go/models';

type Props = {};

const WorkDay = (props: Props) => {
  const {
    appState: { selected },
  } = useAppStore();

  const [worklog, setWorklog] = useState<worklog.WorkLogResult[]>([]);

  useEffect(() => {
    if (selected?.date) {
      const unixTime = moment(selected.date).valueOf();
      GetWorkLog(unixTime).then((res) => {
        setWorklog(res);
      });
    }
  }, [selected]);
  return (
    <div className='py-2'>
      {worklog.map((res) => {
        return res.description;
      })}
    </div>
  );
};

export default WorkDay;
