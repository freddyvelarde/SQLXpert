import { PayloadAction, createSlice } from "@reduxjs/toolkit";

const themeSlice = createSlice({
  name: "theme",
  initialState: {
    value: false,
  },
  reducers: {
    setTheme: (state, actions: PayloadAction<boolean>) => {
      state.value = actions.payload;
    },
  },
});

export const { setTheme } = themeSlice.actions;
export default themeSlice.reducer;
