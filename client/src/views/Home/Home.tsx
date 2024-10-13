import React, { useState } from "react";
import { Card, Carousel, Input } from "../../components";
import { FaSearch } from "react-icons/fa";
import { useNavigate } from "react-router-dom";

const imageUrls = [
  "https://d1csarkz8obe9u.cloudfront.net/posterpreviews/charity-template-banner-design-326a2d18518aadf7713faf82710e7d03_screen.jpg?ts=1662626212",
  "https://d1csarkz8obe9u.cloudfront.net/posterpreviews/charity-template-banner-design-326a2d18518aadf7713faf82710e7d03_screen.jpg?ts=1662626212",
  "https://d1csarkz8obe9u.cloudfront.net/posterpreviews/charity-template-banner-design-326a2d18518aadf7713faf82710e7d03_screen.jpg?ts=1662626212",
];

export function Home() {
  return (
    <div className="w-full flex flex-col px-6 py-6 bg-gray-100 ">
      <div className="w-full flex justify-center">
        <div className="w-[420px]">
          <Input placeholder="جستجوی فراخوان..." icon={<FaSearch />} />
        </div>
      </div>
      <div className="my-6 px-16">
        <Carousel height="400px">
          {imageUrls.map((image, index) => (
            <div key={index} className="h-full">
              <img
                src={image}
                alt={`Slide ${index + 1}`}
                className="object-cover w-full h-full rounded-lg"
                style={{ height: "400px" }}
              />
            </div>
          ))}
        </Carousel>
      </div>
      <div className="px-16 my-12">
        <QuoteOfTheDay
          author="محمد جان"
          quote="شب که میشود به صبح بیچارگان بیاندیش"
        />
      </div>
      <div className="px-16 my-12 ">
        <TrendingRallys />
      </div>
      <div className="px-16 mt-12 ">
        <CardList />
      </div>
    </div>
  );
}

interface QuoteOfTheDayProps {
  quote: string;
  author: string;
}

const QuoteOfTheDay: React.FC<QuoteOfTheDayProps> = ({ quote, author }) => {
  return (
    <div className="mx-auto bg-gradient-to-r from-gray-50 to-gray-100 rounded-lg shadow-lg p-6 text-center">
      <p className="text-xl italic text-gray-700 mb-4">"{quote}"</p>
      <p className="text-md font-semibold text-gray-900">- {author}</p>
      <div className="mt-4">
        <hr className="border-t-2 border-gray-300 mb-4" />
        <p className="text-sm text-gray-500 font-vazir">سخن روز</p>
      </div>
    </div>
  );
};

const TrendingRallys: React.FC = () => {
  const navigate = useNavigate();

  return (
    <>
      <h2 className="text-3xl mb-8">فراخوان های محبوب</h2>
      <Carousel
        settings={{
          slidesToShow: 4, // Default number of slides shown at once
          slidesToScroll: 1, // Number of slides to scroll
          responsive: [
            {
              breakpoint: 1920, // Adjust for large screens
              settings: {
                slidesToShow: 5,
              },
            },
            {
              breakpoint: 1440, // Adjust for large screens
              settings: {
                slidesToShow: 4,
              },
            },
            {
              breakpoint: 1024, // Adjust for large screens
              settings: {
                slidesToShow: 3,
              },
            },
            {
              breakpoint: 768, // Adjust for medium screens (tablets)
              settings: {
                slidesToShow: 2,
              },
            },
            {
              breakpoint: 480, // Adjust for small screens (mobile)
              settings: {
                slidesToShow: 1,
              },
            },
          ],
        }}
        width="100%"
        height="500px"
      >
        {Array.from({ length: 12 }, (_, index) => (
          <div key={index} className="p-2">
            <Card
              imageUrl="https://via.placeholder.com/400x300"
              altText={`Card ${index + 1}`}
            >
              <h2 className="text-2xl font-bold mb-2 text-gray-900">
                Stunning {index + 1}
              </h2>
              <p className="text-gray-700 text-sm">
                Experience the breathtaking beauty of nature with this serene
                landscape. Perfect for unwinding after a long day.
              </p>
              <button
                onClick={() => {
                  navigate(`/rallydetail/2211`);
                }}
                className="mt-4 w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-2 px-4 rounded-lg shadow-md transition-transform transform hover:scale-105"
              >
                Learn More
              </button>
            </Card>
          </div>
        ))}
      </Carousel>
    </>
  );
};

const CardList: React.FC = () => {
  const cards = Array.from({ length: 80 }, (_, i) => ({
    id: i,
    title: `Card Title ${i + 1}`,
    description: `Some description about card ${i + 1}.`,
    imageUrl: "https://via.placeholder.com/400x300",
  }));

  const [currentPage, setCurrentPage] = useState(1);
  const cardsPerPage = 12;

  // Calculate current cards
  const indexOfLastCard = currentPage * cardsPerPage;
  const indexOfFirstCard = indexOfLastCard - cardsPerPage;
  const currentCards = cards.slice(indexOfFirstCard, indexOfLastCard);

  const totalPages = Math.ceil(cards.length / cardsPerPage);

  const handlePageChange = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  const navigate = useNavigate();

  return (
    <div className="min-h-screen">
      <h2 className="text-3xl mb-8">آخرین فراخوان ها</h2>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        {currentCards.map((card) => (
          <Card key={card.id} imageUrl={card.imageUrl} altText={card.title}>
            <h2 className="text-2xl font-bold mb-2 text-gray-900">
              {card.title}
            </h2>
            <p className="text-gray-700 text-sm">{card.description}</p>
            <button
              onClick={() => {
                navigate(`/rallydetail/2211`);
              }}
              className="mt-4 w-full bg-indigo-600 hover:bg-indigo-700 text-white font-semibold py-2 px-4 rounded-lg shadow-md transition-transform transform hover:scale-105"
            >
              Learn More
            </button>
          </Card>
        ))}
      </div>
      {/* Pagination Component */}
      <div className="flex justify-center mt-8">
        <nav>
          <ul className="flex">
            {Array.from({ length: totalPages }, (_, index) => (
              <li className="mx-2" key={index + 1}>
                <button
                  onClick={() => handlePageChange(index + 1)}
                  className={`px-4 py-2 rounded-lg ${
                    currentPage === index + 1
                      ? "bg-indigo-600 text-white"
                      : "bg-gray-200 text-gray-800"
                  } transition duration-300 hover:bg-indigo-500 hover:text-white`}
                >
                  {index + 1}
                </button>
              </li>
            ))}
          </ul>
        </nav>
      </div>
    </div>
  );
};

export default CardList;
