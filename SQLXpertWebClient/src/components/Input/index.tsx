import { DbConnection } from "../../interfaces/dbConnectionConfig";

interface InputProps {
  type: string;
  placeholder: string;
  label: string;
  name: keyof DbConnection;
  state: DbConnection;
  setState: React.Dispatch<React.SetStateAction<DbConnection>>;
}

export default function Input({
  type,
  placeholder,
  label,
  setState,
  state,
  name,
}: InputProps) {
  const handleOnChangeEvent = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setState((prevConfig) => ({
      ...prevConfig,
      [name]: value,
    }));
  };

  return (
    <div>
      <label htmlFor={"formId-" + placeholder}>
        {label}
        <input
          id={"formId-" + placeholder}
          type={type}
          name={name}
          value={state[name]}
          onChange={handleOnChangeEvent}
          placeholder={placeholder}
        />
      </label>
    </div>
  );
}
