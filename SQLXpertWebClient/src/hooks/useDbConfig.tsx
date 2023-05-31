import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../redux/store";
import { setDbConfig } from "../redux/features/dbConfigSlice";
import DbConnection from "../interfaces/dbConnectionConfig";

export default function useDbConfig() {
  const dbConfigConnection = useSelector(
    (state: RootState) => state.dbConfig.value
  );
  const dispatch = useDispatch();

  const storeDbConfig = (dbConfig: DbConnection) => {
    console.log("storing db cofig...");
    dispatch(setDbConfig(dbConfig));
  };

  return {
    storeDbConfig,
    dbConfigConnection,
  };
}
