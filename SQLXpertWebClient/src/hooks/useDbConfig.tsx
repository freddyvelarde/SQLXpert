import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../redux/store";
import { setDbConfig } from "../redux/features/dbConfigSlice";
import DbConnection from "../interfaces/dbConnectionConfig";

export default function useDbConfig() {
  const dbConfig = useSelector((state: RootState) => state.dbConfig.value);
  const dispatch = useDispatch();

  const storeDbConfig = (dbConfig: DbConnection) => {
    dispatch(setDbConfig(dbConfig));
  };

  return {
    storeDbConfig,
    dbConfig
  };
}
