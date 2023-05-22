import { useEffect, useState } from "react";
import Input from "../../components/Input";
import { useNavigate } from "react-router-dom";
import { dbConnection } from "../../data/endpoints";

export default function Connection() {
  const navigate = useNavigate();

  // database config
  const [host, setHost] = useState<string>("localhost");
  const [port, setPort] = useState<string>("5432");
  const [dbName, setDbName] = useState<string>("freddy_db");
  const [user, setUser] = useState<string>("admin");
  const [password, setPassword] = useState<string>("admin");

  // Rest Api response
  const [response, setResponse] = useState<any>([]);

  const navigateToDashboard = async () => {
    const res = await response.connected;
    if (res) {
      navigate(`/${dbName}`, { replace: true });
    }
  };

  const makeQuery = async () => {
    const request = await fetch(dbConnection, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        host,
        port: +port,
        user,
        password,
        dbName,
      }),
    });
    const response = await request.json();
    setResponse(response);
  };

  useEffect(() => {
    navigateToDashboard();
  }, [makeQuery]);

  const handleOnSubmitEvent = (e: React.FormEvent) => {
    e.preventDefault();
    makeQuery();
  };

  return (
    <div>
      <h1>PostgreSql connection</h1>
      <form action="" onSubmit={handleOnSubmitEvent}>
        <Input
          type="text"
          placeholder="username"
          label="USER: "
          state={user}
          setState={setUser}
        />
        <Input
          type="text"
          placeholder="localhost"
          label="HOST: "
          state={host}
          setState={setHost}
        />
        <Input
          type="text"
          placeholder="5432"
          label="PORT: "
          state={port}
          setState={setPort}
        />
        <Input
          type="text"
          placeholder="postgres"
          label="DB NAME: "
          state={dbName}
          setState={setDbName}
        />
        <Input
          type="text"
          placeholder="admin"
          label="DB PASSWORD: "
          state={password}
          setState={setPassword}
        />
        <button>Connect database</button>
      </form>
    </div>
  );
}
