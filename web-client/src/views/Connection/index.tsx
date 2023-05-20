import { useState } from "react";
import Input from "../../components/Input";

export default function Connection() {
  const [host, setHost] = useState<string>("localhost");
  const [port, setPort] = useState<number>(5432);
  const [dbName, setDbName] = useState<string>("freddy_db");
  const [user, setUser] = useState<string>("admin");
  const [quey, setQuery] = useState<string>('select * from "user"');
  const [password, setPassword] = useState<string>("admin");

  const [res, setRes] = useState<[]>([]);

  const makeQuery = async () => {
    const request = await fetch("http://172.19.0.1:7676/query", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        host,
        port,
        user,
        password,
        dbName,
        query: quey,
      }),
    });
    const response = await request.json();
    setRes(response);
  };

  const handleOnSubmitEvent = (e: React.FormEvent) => {
    e.preventDefault();
    makeQuery();
  };

  return (
    <div>
      <h1>connection view</h1>
      <form action="" onSubmit={handleOnSubmitEvent}>
        <Input
          type="text"
          placeholder="SELECT * FROM tablename"
          label="Query: "
          state={quey}
          setState={setQuery}
        />
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
        <button>send data</button>
      </form>
      <button onClick={() => console.log(res)}>getData</button>
    </div>
  );
}
