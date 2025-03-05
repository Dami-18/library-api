import React, { useState, useEffect } from 'react';
import { BACKEND_URL } from './url';

interface Book { // this should match with the struct definition of Book model in backend to directly use this
  CreatedAt: string
  DeletedAt: string
  ID: string;
  UpdatedAt: string
  author: string;
  genre: string; // these must be exactly matching with the json body returned by api response
  rating: number;
  title: string;
}


const Home: React.FC = () => {
  const [title, setTitle] = useState('');
  const [author, setAuthor] = useState('');
  const [rating, setRating] = useState('');
  const [genre, setGenre] = useState('');
  const [filterRating, setFilterRating] = useState('');
  const [filterGenre, setFilterGenre] = useState('');
  const [searchId, setSearchId] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [books, setBooks] = useState<Book[]>([]); // initially empty
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [isLoading, setIsLoading] = useState(false);

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
          'Authorization': `${token}` // try removing bearer
        },
        body: JSON.stringify({ title: title, author: author, genre: genre, rating: rating })
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

  const handleFilter = async () => {
    setIsLoading(true);
    try {

      const queryParams = new URLSearchParams();
      if (filterRating) queryParams.append('rating', filterRating);
      if (filterGenre) queryParams.append('genre', filterGenre);
      queryParams.append('page', currentPage.toString());
      queryParams.append('limit', '10');

      const response = await fetch(`${BACKEND_URL}/books?${queryParams.toString()}`);

      if (!response.ok) {
        throw new Error('Failed to fetch books');
      }
      const data = await response.json();
      setBooks(data.data); 
      // console.log(data.data)
      setTotalPages(Math.ceil(data.total / parseInt(data.limit)));
      setCurrentPage(parseInt(data.page) || 1);

    } catch (error) {
      console.error('Error fetching books:', error);
      setError('Failed to fetch books. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (filterRating || filterGenre) {
      handleFilter();
    }
  }, [currentPage]);

  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
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

      {/* Display of filtered books */}
      {isLoading ? (
        <p>Loading...</p>
      ) : books.length > 0 ? (
        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-2xl font-semibold mb-4">Filtered Books</h2>

          <table className="w-full border-collapse border border-gray-300">
            <thead>
              <tr className="bg-gray-100">
                <th className="border border-gray-300 p-2">ID</th>
                <th className="border border-gray-300 p-2">Title</th>
                <th className="border border-gray-300 p-2">Author</th>
                <th className="border border-gray-300 p-2">Genre</th>
                <th className="border border-gray-300 p-2">Rating</th>
              </tr>
            </thead>
            <tbody>
              {books.map((book) => (
                <tr key={book.ID}>
                  <td className="border border-gray-300 p-2">{book.ID}</td>
                  <td className="border border-gray-300 p-2">{book.title}</td>
                  <td className="border border-gray-300 p-2">{book.author}</td>
                  <td className="border border-gray-300 p-2">{book.genre}</td>
                  <td className="border border-gray-300 p-2">{book.rating}</td>
                </tr>
              ))}
            </tbody>
          </table>

          <div className="mt-4 flex justify-center">
            <button
              onClick={() => setCurrentPage((prev) => Math.max(prev - 1, 1))}
              disabled={currentPage === 1}
              className="bg-purple-500 text-white p-2 rounded hover:bg-purple-600 mr-2"
            >
              Previous
            </button>
            <span className="p-2">
              Page {currentPage} of {totalPages}
            </span>
            <button
              onClick={() => setCurrentPage((prev) => Math.min(prev + 1, totalPages))}
              disabled={currentPage === totalPages}
              className="bg-purple-500 text-white p-2 rounded hover:bg-purple-600 ml-2"
            >
              Next
            </button>
          </div>
        </div>
      ) : null}


    </div>
  );
};

export default Home;