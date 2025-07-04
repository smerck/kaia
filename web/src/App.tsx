import React, { useState } from 'react';
import ReactMarkdown from 'react-markdown';

interface Message {
  sender: 'user' | 'agent';
  text: string;
  backend?: string;
}

const API_URL = '/api/v1/ask';

function App() {
  const [messages, setMessages] = useState<Message[]>([]);
  const [input, setInput] = useState('');
  const [loading, setLoading] = useState(false);
  const [backend, setBackend] = useState<'openai' | 'claude'>('openai');

  const sendMessage = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!input.trim()) return;
    const userMsg: Message = { sender: 'user', text: input, backend };
    setMessages((msgs) => [...msgs, userMsg]);
    setLoading(true);
    try {
      const res = await fetch(API_URL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ prompt: input, backend }),
      });
      const data = await res.json();
      setMessages((msgs) => [
        ...msgs,
        { sender: 'agent', text: data.response || 'No response', backend },
      ]);
    } catch (err) {
      setMessages((msgs) => [
        ...msgs,
        { sender: 'agent', text: 'Error contacting API.', backend },
      ]);
    }
    setInput('');
    setLoading(false);
  };

  const renderMessage = (msg: Message) => {
    if (msg.sender === 'user') {
      return <span>{msg.text}</span>;
    } else {
      return <ReactMarkdown>{msg.text}</ReactMarkdown>;
    }
  };

  return (
    <div style={{ maxWidth: 600, margin: '40px auto', fontFamily: 'sans-serif' }}>
      <h2>kaia: Kubernetes AI Agent</h2>
      
      {/* Backend Toggle */}
      <div style={{ marginBottom: 16, display: 'flex', alignItems: 'center', gap: 8 }}>
        <span>Backend:</span>
        <label style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
          <input
            type="radio"
            name="backend"
            value="openai"
            checked={backend === 'openai'}
            onChange={(e) => setBackend(e.target.value as 'openai' | 'claude')}
          />
          OpenAI
        </label>
        <label style={{ display: 'flex', alignItems: 'center', gap: 4 }}>
          <input
            type="radio"
            name="backend"
            value="claude"
            checked={backend === 'claude'}
            onChange={(e) => setBackend(e.target.value as 'openai' | 'claude')}
          />
          Claude
        </label>
      </div>

      <div style={{ border: '1px solid #ccc', borderRadius: 8, padding: 16, minHeight: 300, background: '#fafbfc' }}>
        {messages.length === 0 && <div style={{ color: '#888' }}>Ask me about your Kubernetes environment!</div>}
        {messages.map((msg, i) => (
          <div key={i} style={{ textAlign: msg.sender === 'user' ? 'right' : 'left', margin: '8px 0' }}>
            <div style={{
              display: 'inline-block',
              background: msg.sender === 'user' ? '#d1e7dd' : '#e7eaf6',
              color: '#222',
              borderRadius: 16,
              padding: '8px 16px',
              maxWidth: '80%',
              textAlign: 'left',
            }}>
              {renderMessage(msg)}
            </div>
            {msg.backend && (
              <div style={{ fontSize: '0.8em', color: '#666', marginTop: 4 }}>
                via {msg.backend}
              </div>
            )}
          </div>
        ))}
        {loading && <div style={{ color: '#888' }}>Agent is thinking...</div>}
      </div>
      <form onSubmit={sendMessage} style={{ marginTop: 16, display: 'flex', gap: 8 }}>
        <input
          type="text"
          value={input}
          onChange={e => setInput(e.target.value)}
          placeholder="Type your question..."
          style={{ flex: 1, padding: 8, borderRadius: 8, border: '1px solid #ccc' }}
          disabled={loading}
        />
        <button type="submit" disabled={loading || !input.trim()} style={{ padding: '8px 16px', borderRadius: 8, background: '#1976d2', color: '#fff', border: 'none' }}>
          Send
        </button>
      </form>
    </div>
  );
}

export default App;
