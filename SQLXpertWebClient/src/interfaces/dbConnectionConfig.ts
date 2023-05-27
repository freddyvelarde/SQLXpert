export default interface DbConnection {
  host: string;
  port: string;
  dbName: string;
  user: string;
  password: string;
  workspace?: string;
  query?: string;
}
