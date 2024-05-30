/* eslint-disable @typescript-eslint/no-unused-vars */
import React, {
  Dispatch,
  ReactNode,
  useContext,
  useEffect,
  useReducer,
} from 'react';
import { appstore, worklog } from '../wailsjs/go/models';
import { GetUserConfig } from '../wailsjs/go/main/App';

interface IAppState {
  api?: appstore.UserConfig;
  selected?: worklog.WorkDay;
}
interface IApp {
  appState: IAppState;
  dispatch: React.Dispatch<Action>;
}

const AppContext = React.createContext<IApp>({
  appState: {},
  dispatch: () => {},
});

const initialState: IAppState = {};

type Action =
  | {
      type: 'SET_USER_CONFIG';
      payload: appstore.UserConfig;
    }
  | { type: 'SET_SELECTED_DATE'; payload: worklog.WorkDay };

const reducer = (state: IAppState, action: Action): IAppState => {
  switch (action.type) {
    case 'SET_USER_CONFIG':
      return { ...state, api: action.payload };
      break;
    case 'SET_SELECTED_DATE':
      return { ...state, selected: action.payload };
      break;

    default:
      return state;
      break;
  }
};

const AppContextProvider = ({
  children,
}: // model,
{
  children: React.ReactNode;
  // model: ProductVariantType;
}): JSX.Element => {
  const [state, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    GetUserConfig()
      .then((userConfig) => {
        dispatch({ type: 'SET_USER_CONFIG', payload: userConfig });
      })
      .catch((config) => {});
  }, []);

  return (
    <AppContext.Provider value={{ appState: state, dispatch: dispatch }}>
      {children}
    </AppContext.Provider>
  );
};

export const useAppStore = (): IApp => useContext(AppContext);
export default AppContextProvider;
