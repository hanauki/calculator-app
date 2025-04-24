import { useEffect, useState } from 'react';

export default async function handler(req, res) {
  if (req.method === 'POST') {
    const { expression } = req.body;

    const result = await rpcCall(expression); // Use RPC to call Go service

    res.status(200).json({ result });
  }
}

async function rpcCall(expression) {
  // Placeholder for RPC communication
  return 'calculated result';
}
