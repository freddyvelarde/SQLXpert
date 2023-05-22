interface InputProps {
  type: string;
  placeholder: string;
  label: string;
  state: string | number;
  setState: React.Dispatch<React.SetStateAction<any>>;
}

export default function Input({
  type,
  placeholder,
  label,
  setState,
  state,
}: InputProps) {
  const handleOnChangeEvent = (e: React.ChangeEvent<HTMLInputElement>) => {
    setState(e.target.value);
  };

  return (
    <div>
      <label htmlFor={"formId-" + placeholder}>
        {label}
        <input
          id={"formId-" + placeholder}
          type={type}
          value={state}
          onChange={handleOnChangeEvent}
          placeholder={placeholder}
        />
      </label>
    </div>
  );
}
