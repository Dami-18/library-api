import React, { useState } from 'react';
import { BACKEND_URL } from './url';

const Home: React.FC = () => {
  const [title, setTitle] = useState('');
  const [author, setAuthor] = useState('');
  const [rating, setRating] = useState('');
  const [genre, setGenre] = useState('');
  const [filterRating, setFilterRating] = useState('');
  const [filterGenre, setFilterGenre] = useState('');
  const [searchId, setSearchId] = useState('');
  const [error, setError] = useState<string | null>(null);

  const handleAddBook = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    const token = localStorage.getItem('token');
    if (!token) {
      setError('You must be logged in to add a book.');
      return;
    }

    try {
      const response = await fetch(`${BACKEND_URL}/books`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({ title, author, genre , rating: Number(rating)})
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to add book');
      }

      const data = await response.json();
      console.log('Book added successfully:', data);
      
      // Clear the form
      setTitle('');
      setAuthor('');
      setRating('');
      setGenre('');

    } catch (err: any) {
      console.error('Error adding book:', err.message);
      setError(err.message);
    }
  };

  const handleFilter = () => {
    // Filter books logic here
    console.log('Filtering books:', { rating: filterRating, genre: filterGenre });
  };

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    // Search book by ID logic here
    console.log('Searching for book with ID:', searchId);
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold text-center mb-8">Book Library Home</h1>

      {/* Add Book Form */}
      <div className="mb-8 bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-2xl font-semibold mb-4">Add a New Book</h2>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <form onSubmit={handleAddBook} className="space-y-4">
          <input
            type="text"
            placeholder="Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-full p-2 border rounded"
          />
          <input
            type="text"
            placeholder="Author"
            value={author}
            onChange={(e) => setAuthor(e.target.value)}
            className="w-full p-2 border rounded"
          />

          <input
            type="text"
            placeholder="Genre"
            value={genre}
            onChange={(e) => setGenre(e.target.value)}
            className="w-full p-2 border rounded"
          />

          <input
            type="number"
            placeholder="Rating (1-5)"
            value={rating}
            onChange={(e) => setRating(e.target.value)}
            min="0"
            max="5"
            step="0.1"
            className="w-full p-2 border rounded"
          />
          
          <button type="submit" className="w-full bg-purple-500 text-white p-2 rounded hover:bg-purple-600">
            Add Book
          </button>
        </form>
      </div>

      {/* Filter Books */}
      <div className="mb-8 bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-2xl font-semibold mb-4">Filter Books</h2>
        <div className="flex space-x-4">
          <input
            type="text"
            placeholder="Filter by Rating"
            value={filterRating}
            onChange={(e) => setFilterRating(e.target.value)}
            className="p-2 border rounded"
          >
          </input>

          <input
            type="text"
            placeholder="Filter by Genre"
            value={filterGenre}
            onChange={(e) => setFilterGenre(e.target.value)}
            className="p-2 border rounded flex-grow"
          />

          <button onClick={handleFilter} className="bg-purple-500 text-white p-2 rounded hover:bg-purple-600">
            Apply Filters
          </button>
        </div>
      </div>

      {/* Search Book by ID */}
      <div className="bg-white p-6 rounded-lg shadow-md">
        <h2 className="text-2xl font-semibold mb-4">Search Book by ID</h2>
        <form onSubmit={handleSearch} className="flex space-x-4">
          <input
            type="text"
            placeholder="Enter Book ID"
            value={searchId}
            onChange={(e) => setSearchId(e.target.value)}
            className="p-2 border rounded flex-grow"
          />
          <button type="submit" className="bg-purple-500 text-white p-2 rounded hover:bg-purple-600">
            Search
          </button>
        </form>
      </div>
    </div>
  );
};

export default Home;