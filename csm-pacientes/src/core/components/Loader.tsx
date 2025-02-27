import { Component } from "solid-js";

const Loader: Component<{ size?: string, fullScreen?: boolean }> = (props) => {
	return (
		<div class={props.fullScreen ? "loader-overlay" : "loader-wrapper"}>
			<div class="loader" style={{ width: props.size || "50px", height: props.size || "50px" }}></div>
		</div>
	);
};

export default Loader;
