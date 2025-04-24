import { render, screen, fireEvent } from '@testing-library/react';
import Calculator from '../components/Calculator';

test('renders calculator UI', () => {
  render(<Calculator />);
  const button1 = screen.getByText('1');
  fireEvent.click(button1);
  const result = screen.getByText(/result/i);
  expect(result).toBeInTheDocument();
});
