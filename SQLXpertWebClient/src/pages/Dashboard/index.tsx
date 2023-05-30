import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import DbConnection from "../../interfaces/dbConnectionConfig";
import useHttpRequest from "../../hooks/useHttpRequest";
import { query } from "../../data/endpoints";
import useDbConfig from "../../hooks/useDbConfig";

export default function Dashboard() {
  const { paramsId } = useParams();
  const { dbCofigConnection } = useDbConfig();

  const [text, setText] = useState("");

  const { data, fetchData } = useHttpRequest<DbConnection>(query, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      host: dbCofigConnection.host,
      port: +dbCofigConnection.port,
      user: dbCofigConnection.user,
      password: dbCofigConnection.password,
      dbName: dbCofigConnection.dbName,
      query: text,
    }),
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
      <button onClick={() => console.log(data)}>data</button>
    </>
  );
}
