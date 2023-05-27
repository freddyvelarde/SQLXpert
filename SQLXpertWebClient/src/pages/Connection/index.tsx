import { useEffect, useState } from "react";
import Input from "../../components/Input";
import { useNavigate } from "react-router-dom";
import { dbConnection } from "../../data/endpoints";
import DbConnection from "../../interfaces/dbConnectionConfig";
import useHttpRequest from "../../hooks/useHttpRequest";
import useDatabases from "../../hooks/useDatabases";
import { ConnectionResponse } from "../../interfaces/HttpResponses";
import { emptySpaceValidation } from "../../utils/stringValidation";
// import useDbConfig from "../../hooks/useDbConfig";

export default function Connection() {
  const { addNewDatabase } = useDatabases();
  const navigate = useNavigate();

  // const { dbConfig, storeDbConfig } = useDbConfig();

  // database config
  const [dbConfig, storeDbConfig] = useState<DbConnection>({
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
      host: dbConfig.host,
      port: +dbConfig.port,
      user: dbConfig.user,
      password: dbConfig.password,
      dbName: dbConfig.dbName,
    }),
  });

  const navigateToDashboard = () => {
    const res = data?.connected;
    if (res) {
      navigate(`/${dbConfig.workspace}`, { replace: true });
    }
  };

  useEffect(() => {
    navigateToDashboard();
  }, [fetchData]);

  const handleOnSubmitEvent = (e: React.FormEvent) => {
    e.preventDefault();

    if (!dbConfig.workspace) return;
    if (!emptySpaceValidation(dbConfig.workspace))
      return console.log("no empty space");

    fetchData();

    addNewDatabase(dbConfig);
  };

  return (
    <div>
      <h1>PostgreSql connection</h1>
      <form action="" onSubmit={handleOnSubmitEvent}>
        <Input
          type="text"
          placeholder="Workspace"
          label="WORKSPACE: "
          state={dbConfig}
          name="workspace"
          setState={storeDbConfig}
        />
        <Input
          type="text"
          placeholder="username"
          label="USER: "
          state={dbConfig}
          name="user"
          setState={storeDbConfig}
        />
        <Input
          type="text"
          placeholder="localhost"
          label="HOST: "
          name="host"
          state={dbConfig}
          setState={storeDbConfig}
        />
        <Input
          type="text"
          placeholder="5432"
          label="PORT: "
          name="port"
          state={dbConfig}
          setState={storeDbConfig}
        />
        <Input
          type="text"
          placeholder="postgres"
          label="DB NAME: "
          name="dbName"
          state={dbConfig}
          setState={storeDbConfig}
        />
        <Input
          type="text"
          placeholder="admin"
          label="DB PASSWORD: "
          name="password"
          state={dbConfig}
          setState={storeDbConfig}
        />
        <button>Connect database</button>
      </form>
    </div>
  );
}
