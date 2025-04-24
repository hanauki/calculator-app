import { useState } from 'react';

export default function Calculator() {
  const [input, setInput] = useState('');
  const [result, setResult] = useState('');

  const handleClick = (value) => {
    setInput(input + value);
  };

  const handleCalculate = async () => {
    const response = await fetch('/api/calculate', {
      method: 'POST',
      body: JSON.stringify({ expression: input }),
      headers: { 'Content-Type': 'application/json' },
    });
    const data = await response.json();
    setResult(data.result);
  };

  return (
    <div>
      <input type="text" value={input} readOnly />
      <div>
        <button onClick={() => handleClick('1')}>1</button>
        <button onClick={() => handleClick('2')}>2</button>
        <button onClick={() => handleClick('+')}>+</button>
        {/* Add more buttons for digits and operators */}
      </div>
      <button onClick={handleCalculate}>Calculate</button>
      <div>Result: {result}</div>
    </div>
  );
}
