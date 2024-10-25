import { createSignal } from 'solid-js';
import bellIcon from '../assets/bell.png';

interface NotificationProps {
  title: string;
  message: string;
  time: string;
}

const NotificationIcon = (props: NotificationProps) => {
  const [isExpanded, setIsExpanded] = createSignal(false);

  const toggleNotifications = () => {
    setIsExpanded(!isExpanded());
  };

  return (
    <div class="notification-container">
      <img
        src={bellIcon}
        alt="Notificaciones"
        class="notification-icon"
        onClick={toggleNotifications}
      />
      {isExpanded() && (
        <div class="notification-panel">
          <h4>{props.title}</h4>
          <p>{props.message}</p>
          <span>{props.time}</span>
        </div>
      )}
    </div>
  );
};

export default NotificationIcon;
