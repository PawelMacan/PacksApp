'use client'

import { useState } from 'react'

export default function Home() {
  const [amount, setAmount] = useState('')
  const [result, setResult] = useState<any>(null)
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setResult(null)

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/calculate`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ amount: parseInt(amount) }),
      })

      const data = await res.json()
      setResult(data)
    } catch (err) {
      console.error('API error:', err)
      alert('Wystąpił błąd przy pobieraniu danych z API')
    } finally {
      setLoading(false)
    }
  }

  return (
      <main className="min-h-screen bg-gray-100 flex flex-col items-center justify-start p-6">
        <div className="w-full max-w-xl bg-white shadow-md rounded-lg p-6 mt-10">
          <h1 className="text-2xl font-semibold mb-4">🧮 Obliczanie Paczek</h1>

          <form onSubmit={handleSubmit} className="flex gap-4 mb-6">
            <input
                type="number"
                value={amount}
                onChange={(e) => setAmount(e.target.value)}
                placeholder="Ilość sztuk"
                min={1}
                required
                className="flex-1 border border-gray-300 rounded px-4 py-2"
            />
            <button
                type="submit"
                disabled={loading}
                className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded"
            >
              {loading ? 'Liczenie...' : 'Oblicz'}
            </button>
          </form>

          {result && (
              <div>
                <h2 className="text-xl font-semibold mb-2">📦 Wynik:</h2>
                <div className="bg-gray-50 border rounded p-4 text-sm">
                  <p><strong>Zażądano:</strong> {result.requested_amount}</p>
                  <p><strong>Dostarczono:</strong> {result.total_items}</p>
                  <p><strong>Nadwyżka:</strong> {result.overage}</p>
                  <p><strong>Liczba paczek:</strong> {result.total_packs}</p>

                  <div className="mt-2">
                    <strong>Konfiguracja paczek:</strong>
                    <ul className="list-disc list-inside">
                      {Object.entries(result.packs as Record<string, number>).map(([size, count]) => (

                          <li key={size}>
                            {count} × {size}
                          </li>
                      ))}
                    </ul>
                  </div>
                </div>
              </div>
          )}
        </div>
      </main>
  )
}
