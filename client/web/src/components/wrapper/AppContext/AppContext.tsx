// AppContext.tsx
import React, { createContext, useContext, useReducer, ReactNode } from "react";
import { AppState, AppActions, User, Rally, Charity } from "./AppContext.types";

// Initial State
const initialState: AppState = {
  user: null,
  rallyList: [],
  selectedRally: null,
  charity: null,
};

// Action Types
type Action =
  | { type: "SET_USER"; payload: User | null }
  | { type: "SET_APPEALS_LIST"; payload: Rally[] }
  | { type: "SET_SELECTED_APPEAL"; payload: Rally | null }
  | { type: "SET_CHARITY"; payload: Charity | null };

// Reducer Function
const appReducer = (state: AppState, action: Action): AppState => {
  switch (action.type) {
    case "SET_USER":
      return { ...state, user: action.payload };
    case "SET_APPEALS_LIST":
      return { ...state, rallyList: action.payload };
    case "SET_SELECTED_APPEAL":
      return { ...state, selectedRally: action.payload };
    case "SET_CHARITY":
      return { ...state, charity: action.payload };
    default:
      return state;
  }
};

// Create Context
const AppStateContext = createContext<AppState | undefined>(undefined);
const AppActionsContext = createContext<AppActions | undefined>(undefined);

// Provider Component
interface AppProviderProps {
  children: ReactNode;
}

export const AppProvider = ({ children }: AppProviderProps) => {
  const [state, dispatch] = useReducer(appReducer, initialState);

  const actions: AppActions = {
    setUser: (user) => dispatch({ type: "SET_USER", payload: user }),
    setRallyList: (appeals) =>
      dispatch({ type: "SET_APPEALS_LIST", payload: appeals }),
    setSelectedRally: (appeal) =>
      dispatch({ type: "SET_SELECTED_APPEAL", payload: appeal }),
    setCharity: (charity) =>
      dispatch({ type: "SET_CHARITY", payload: charity }),
  };

  return (
    <AppStateContext.Provider value={state}>
      <AppActionsContext.Provider value={actions}>
        {children}
      </AppActionsContext.Provider>
    </AppStateContext.Provider>
  );
};

// Custom hooks for state and actions
export const useAppState = () => {
  const context = useContext(AppStateContext);
  if (!context) {
    throw new Error("useAppState must be used within an AppProvider");
  }
  return context;
};

export const useAppActions = () => {
  const context = useContext(AppActionsContext);
  if (!context) {
    throw new Error("useAppActions must be used within an AppProvider");
  }
  return context;
};
