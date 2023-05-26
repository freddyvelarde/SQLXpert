import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../redux/store";
import DbConnection from "../interfaces/dbConnectionConfig";
import { setDatabases } from "../redux/features/databasesSlice";

export default function useDatabases() {
  const databases = useSelector((state: RootState) => state.databases.value);
  const dispatch = useDispatch();

  const setDatabasesIntoRedux = (databases: DbConnection[]) => {
    dispatch(setDatabases(databases));
  };

  const addNewDatabase = (db: DbConnection) => {
    const dbFound = databases.find(
      (dbRepeated: DbConnection) => dbRepeated.workspace == db.workspace
    );
    if (dbFound) {
      return console.log(dbFound);
    }
    dispatch(setDatabases([...databases, db]));
  };

  return {
    addNewDatabase,
    databases,
    setDatabasesIntoRedux,
  };
}
