import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import DbConnection from "../../interfaces/dbConnectionConfig";

const initialState: DbConnection[] = [];

const databasesSlice = createSlice({
  name: "databases",
  initialState: {
    value: initialState,
  },
  reducers: {
    setDatabases: (state, actions: PayloadAction<DbConnection[]>) => {
      state.value = actions.payload;
    },
  },
});

export const { setDatabases } = databasesSlice.actions;
export default databasesSlice.reducer;
