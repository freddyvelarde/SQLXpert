import { useEffect, useState } from "react";
import Input from "../../components/Input";
import { useNavigate } from "react-router-dom";
import { dbConnection } from "../../data/endpoints";
import { DbConnection } from "../../interfaces/dbConnectionConfig";
import useHttpRequest from "../../hooks/useHttpRequest";

interface Response {
  connected: boolean;
}

export default function Connection() {
  const navigate = useNavigate();

  // database config
  const [dbConfigConnection, setDbConfigConnection] = useState<DbConnection>({
    host: "localhost",
    port: "5432",
    dbName: "freddy_db",
    user: "admin",
    password: "admin",
  });
  const { data, fetchData } = useHttpRequest<Response>(dbConnection, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      host: dbConfigConnection.host,
      port: +dbConfigConnection.port,
      user: dbConfigConnection.user,
      password: dbConfigConnection.password,
      dbName: dbConfigConnection.dbName,
    }),
  });

  // Rest Api response
  // const [response, setResponse] = useState<any>([]);

  const navigateToDashboard = () => {
    const res = data?.connected;
    if (res) {
      navigate(`/${dbConfigConnection.dbName}`, { replace: true });
    }
  };

  // const makeQuery = async () => {
  //   const request = await fetch(dbConnection, {
  //     method: "POST",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //     body: JSON.stringify({
  //       host: dbConfigConnection.host,
  //       port: +dbConfigConnection.port,
  //       user: dbConfigConnection.user,
  //       password: dbConfigConnection.password,
  //       dbName: dbConfigConnection.dbName,
  //     }),
  //   });
  //   const response = await request.json();
  //   setResponse(response);
  // };

  useEffect(() => {
    navigateToDashboard();
  }, [fetchData]);

  const handleOnSubmitEvent = (e: React.FormEvent) => {
    e.preventDefault();
    fetchData();
  };

  return (
    <div>
      <h1>PostgreSql connection</h1>
      <form action="" onSubmit={handleOnSubmitEvent}>
        <Input
          type="text"
          placeholder="username"
          label="USER: "
          state={dbConfigConnection}
          name="user"
          setState={setDbConfigConnection}
        />
        <Input
          type="text"
          placeholder="localhost"
          label="HOST: "
          name="host"
          state={dbConfigConnection}
          setState={setDbConfigConnection}
        />
        <Input
          type="text"
          placeholder="5432"
          label="PORT: "
          name="port"
          state={dbConfigConnection}
          setState={setDbConfigConnection}
        />
        <Input
          type="text"
          placeholder="postgres"
          label="DB NAME: "
          name="dbName"
          state={dbConfigConnection}
          setState={setDbConfigConnection}
        />
        <Input
          type="text"
          placeholder="admin"
          label="DB PASSWORD: "
          name="password"
          state={dbConfigConnection}
          setState={setDbConfigConnection}
        />
        <button>Connect database</button>
      </form>
    </div>
  );
}