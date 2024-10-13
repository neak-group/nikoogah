// types.ts
export interface User {
  id: string;
  name: string;
  email: string;
}

export interface Rally {
  id: string;
  title: string;
  description: string;
  status: string;
}

export interface Charity {
  id: string;
  name: string;
}

export interface AppState {
  user: User | null;
  rallyList: Rally[];
  selectedRally: Rally | null;
  charity: Charity | null;
}

export interface AppActions {
  setUser: (user: User | null) => void;
  setRallyList: (rallyes: Rally[]) => void;
  setSelectedRally: (rally: Rally | null) => void;
  setCharity: (charity: Charity | null) => void;
}
