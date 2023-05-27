import { configureStore } from "@reduxjs/toolkit";
import theme from "./features/themeSlices";
import databases from "./features/databasesSlice";
import dbConfig from "./features/dbConfigSlice.ts";

export const store = configureStore({
  reducer: {
    theme,
    databases,
    dbConfig,
  },
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
// Inferred type: {posts: PostsState, comments: CommentsState, users: UsersState}
export type AppDispatch = typeof store.dispatch;
