interface MinimalistCardProps {
  image: string;
  title: string;
  content: string;
}

const MinimalistCard = ({ image, title, content }: MinimalistCardProps) => {
    return (
      <div className="minimalist-card">
        <img src={image} alt={title} className="card-image" />
        <div className="card-content">
          <h2>{title}</h2>
          <p>{content}</p>
        </div>
      </div>
    );
  }
  
  export default MinimalistCard;