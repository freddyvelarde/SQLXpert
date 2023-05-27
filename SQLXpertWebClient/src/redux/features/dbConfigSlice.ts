import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import DbConnection from "../../interfaces/dbConnectionConfig";

const initialState: DbConnection = {
  host: "",
  port: "",
  user: "",
  dbName: "",
  password: ""
};

const dbConfig = createSlice({
  name: "databases",
  initialState: {
    value: initialState,
  },
  reducers: {
    setDbConfig: (state, actions: PayloadAction<DbConnection>) => {
      state.value = actions.payload;
    },
  },
});

export const { setDbConfig } = dbConfig.actions;
export default dbConfig.reducer;
