/* eslint-disable @typescript-eslint/no-unused-vars */
import React, {
  Dispatch,
  ReactNode,
  useContext,
  useEffect,
  useReducer,
} from 'react';
import { appstore } from '../wailsjs/go/models';
import { GetUserConfig } from '../wailsjs/go/main/App';

interface IAppState {
  api?: appstore.UserConfig;
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

type Action = {
  type: 'SET_USER_CONFIG';
  payload: appstore.UserConfig;
};

const reducer = (state: IAppState, action: Action): IAppState => {
  if (action.type === 'SET_USER_CONFIG') {
    return {
      ...state,
      api: action.payload,
    };
  }
  return state;
};

const AppContextProvider = ({
  children,
}: // model,
{
  children: React.ReactNode;
  // model: ProductVariantType;
}): JSX.Element => {
  const [state, dispatch] = useReducer(reducer, initialState);

  GetUserConfig()
    .then((userConfig) => {
      dispatch({ type: 'SET_USER_CONFIG', payload: userConfig });
    })
    .catch((config) => {});

  return (
    <AppContext.Provider value={{ appState: state, dispatch: dispatch }}>
      {children}
    </AppContext.Provider>
  );
};

export const useAppStore = (): IApp => useContext(AppContext);
export default AppContextProvider;
