import { useEffect, useState } from "react";
import Input from "../../components/Input";
import { useNavigate } from "react-router-dom";
import { dbConnection } from "../../data/endpoints";
import DbConnection from "../../interfaces/dbConnectionConfig";
import useHttpRequest from "../../hooks/useHttpRequest";
import useDatabases from "../../hooks/useDatabases";
import { ConnectionResponse } from "../../interfaces/HttpResponses";
import { emptySpaceValidation } from "../../utils/stringValidation";
import useDbConfig from "../../hooks/useDbConfig";
import Error from "../../components/Error";

export default function Connection() {
  const { addNewDatabase, databases } = useDatabases();
  const [error, setError] = useState({ failed: false, message: "" });
  const navigate = useNavigate();

  const { storeDbConfig } = useDbConfig();

  // database config
  const [dbConfigForm, setDbCofigForm] = useState<DbConnection>({
    host: "localhost",
    port: "5432",
    dbName: "freddy_db",
    user: "admin",
    password: "admin",
    workspace: "workspace name",
  });

  const { data, fetchData } = useHttpRequest<ConnectionResponse>(dbConnection, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      host: dbConfigForm.host,
      port: +dbConfigForm.port,
      user: dbConfigForm.user,
      password: dbConfigForm.password,
      dbName: dbConfigForm.dbName,
    }),
  });

  const navigateToDashboard = () => {
    const res = data?.connected;
    if (res) {
      navigate(`/database/${dbConfigForm.workspace}`, { replace: true });
    }
  };

  const handleOnSubmitEvent = (e: React.FormEvent) => {
    e.preventDefault();

    if (!dbConfigForm.workspace)
      return setError({
        failed: true,
        message: "Workspace field does not exist.",
      });
    if (!emptySpaceValidation(dbConfigForm.workspace))
      return setError({
        failed: true,
        message: "Please no empty spaces in workspace field.",
      });

    const dbFound = databases.find(
      (dbRepeated: DbConnection) =>
        dbRepeated.workspace == dbConfigForm.workspace
    );
    if (dbFound) {
      `Your workspace name already exist!`;
      return setError({
        failed: true,
        message: `Your workspace name: "${dbConfigForm.workspace}" already exist!`,
      });
    }

    addNewDatabase(dbConfigForm);

    storeDbConfig(dbConfigForm);
    fetchData();
  };

  useEffect(() => {
    navigateToDashboard();
  }, [fetchData]);

  const setPreviousDbConfig = (dbConfig: DbConnection) => {
    // const previousConfig = databases.find(
    //   (db: DbConnection) => db.workspace === workspaceName
    // );
    // if (!previousConfig) return;
    setDbCofigForm(dbConfig);
    storeDbConfig(dbConfig);
    fetchData();
  };

  return (
    <div>
      <h1>PostgreSql connection</h1>
      <form action="" onSubmit={handleOnSubmitEvent}>
        <Input
          type="text"
          placeholder="Workspace"
          label="WORKSPACE: "
          state={dbConfigForm}
          name="workspace"
          setState={setDbCofigForm}
        />
        <Input
          type="text"
          placeholder="username"
          label="USER: "
          state={dbConfigForm}
          name="user"
          setState={setDbCofigForm}
        />
        <Input
          type="text"
          placeholder="localhost"
          label="HOST: "
          name="host"
          state={dbConfigForm}
          setState={setDbCofigForm}
        />
        <Input
          type="text"
          placeholder="5432"
          label="PORT: "
          name="port"
          state={dbConfigForm}
          setState={setDbCofigForm}
        />
        <Input
          type="text"
          placeholder="postgres"
          label="DB NAME: "
          name="dbName"
          state={dbConfigForm}
          setState={setDbCofigForm}
        />
        <Input
          type="text"
          placeholder="admin"
          label="DB PASSWORD: "
          name="password"
          state={dbConfigForm}
          setState={setDbCofigForm}
        />
        <button>Connect database</button>
      </form>
      <button onClick={() => console.log(data)}>Data config</button>

      {error.failed ? <Error message={error.message} /> : ""}

      {databases.map((db) => (
        <button onClick={() => (db.workspace ? setPreviousDbConfig(db) : null)}>
          {db.workspace}
        </button>
      ))}
    </div>
  );
}
