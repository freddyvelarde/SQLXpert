import { useState } from "react";
import { useParams } from "react-router-dom";
import DbConnection from "../../interfaces/dbConnectionConfig";
import useHttpRequest from "../../hooks/useHttpRequest";
import { query } from "../../data/endpoints";
import useDbConfig from "../../hooks/useDbConfig";

export default function Dashboard() {
  const { paramsId } = useParams();
  const { dbConfigConnection } = useDbConfig();

  // useEffect(() => {
  //   if (dbConfigConnection.workspace != paramsId) {
  //
  //   }
  // }, []);

  const [text, setText] = useState("");

  const { data, fetchData } = useHttpRequest<DbConnection>(query, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      host: dbConfigConnection.host,
      port: +dbConfigConnection.port,
      user: dbConfigConnection.user,
      password: dbConfigConnection.password,
      dbName: dbConfigConnection.dbName,
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
