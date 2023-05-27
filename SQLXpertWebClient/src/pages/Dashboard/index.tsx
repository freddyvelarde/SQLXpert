import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import useDatabases from "../../hooks/useDatabases";
import DbConnection from "../../interfaces/dbConnectionConfig";
import useHttpRequest from "../../hooks/useHttpRequest";
import { query } from "../../data/endpoints";

export default function Dashboard() {
  const { paramsId } = useParams();
  const { databases } = useDatabases();
  const [db, setDb] = useState<DbConnection>({
    host: "",
    password: "",
    user: "",
    dbName: "",
    port: "",
  });

  const getItem = () => {
    const item = databases.find((item) => item.workspace === paramsId);
    if (!item) return;
    setDb(item);
  };

  useEffect(() => {
    getItem();
  }, []);

  const [text, setText] = useState("");

  const { data, fetchData } = useHttpRequest<DbConnection>(query, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      host: db.host,
      port: +db.port,
      user: db.user,
      password: db.password,
      dbName: db.dbName,
      query: text,
    }),
    // body: JSON.stringify({
    //   host: "localhost",
    //   port: 5432,
    //   user: "admin",
    //   password: "admin",
    //   dbName: "freddy_db",
    //   query: text,
    // }),
  });

  const handleChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    fetchData();
  };

  return (
    <>
      <h1>Dashboard: {paramsId}</h1>
      <form action="" onSubmit={handleSubmit}>
        <textarea
          rows={10}
          cols={50}
          value={text}
          onChange={handleChange}
        ></textarea>
        <button type="submit">Submit</button>
      </form>
      <button onClick={() => console.log(databases)}>data</button>
    </>
  );
}
