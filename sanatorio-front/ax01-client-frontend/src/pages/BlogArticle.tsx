import Article from "../components/Article";
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom'; 

type BlogPostType = {
    id: number;
    title: string;
    category: string;
    content: string;
    author: string;
    date: string;
  };
  
  // Función para analizar una cadena de fecha y devolver un objeto Date
  function parseDate(dateString: string): Date {
    if (!dateString) {
      return new Date();
    }
    
    const parts = dateString.split(/[^\d]/);
    const year = parseInt(parts[0]);
    const month = parseInt(parts[1]) - 1;
    const day = parseInt(parts[2]);
    const hour = parseInt(parts[3]);
    const minute = parseInt(parts[4]);
    const second = parseInt(parts[5]) || 0;
    return new Date(year, month, day, hour, minute, second);
  }
  
  // Función para formatear una fecha en formato legible
  function formatDate(dateString: string): string {
    const date = parseDate(dateString);
    const formattedDate = `${padZero(date.getDate())}/${padZero(date.getMonth() + 1)}/${date.getFullYear()} ${padZero(date.getHours())}:${padZero(date.getMinutes())}:${padZero(date.getSeconds())}`;
    return formattedDate;
  }
  
  // Función para agregar un cero delante de un número si es menor que 10
  function padZero(value: number): string {
    return value < 10 ? `0${value}` : `${value}`;
  }
  
  // Componente BlogArticle
  const BlogArticle = () => {
    const { id } = useParams<{ id: string }>(); // Obtén el parámetro 'id' de la URL utilizando useParams
  
    const [article, setArticle] = useState<BlogPostType | null>(null); // Estado para almacenar el artículo del blog
  
    // Efecto para cargar el artículo del blog cuando el componente se monta o el 'id' cambia
    useEffect(() => {
      const fetchArticle = async () => {
        try {
          const response = await fetch(`https://api.ax01.dev/post/get/${id}`); // Realiza una solicitud para obtener el artículo del blog
          if (response.ok) {
            const data: BlogPostType = await response.json(); // Convierte la respuesta en JSON
            setArticle(data); // Actualiza el estado con el artículo del blog
          } else {
            console.error("Error al cargar el artículo del blog");
          }
        } catch (error) {
          console.error("Error al cargar el artículo del blog:", error);
        }
      };
  
      fetchArticle(); // Llama a la función para cargar el artículo del blog
    }, [id]); // El efecto se ejecuta cuando cambia el valor de 'id'
  
    // Si el artículo aún no se ha cargado, muestra un mensaje de carga
    if (!article) {
      return <div>Cargando...</div>;
    }
  
    const formattedDate = formatDate(article.date); // Formatea la fecha del artículo
  
    // Renderiza el componente Article con los datos del artículo
    return (
      <Article
        id={article.id}
        title={article.title}
        category={article.category}
        content={article.content}
        author={article.author}
        date={formattedDate}
      />
    );
  };
  
  export default BlogArticle;