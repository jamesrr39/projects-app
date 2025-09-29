import { useGetProjectsListing } from "../../openapi/generated/projectsAppComponents";
import Error from "../Error";
import Loading from "../Loading";

function ProjectListing() {
    const { data, isLoading, error } = useGetProjectsListing({});
    if (error) {
        return <Error error={error} />
    }

    if (isLoading) {
        return <Loading />
    }

    if (!data) {
        return null;
    }

    return (
        <div>
            <h1>Projects</h1>
            <table>
                {data.projects.map((project, idx) => {
                    return (
                        <tr key={idx}>
                            <td>{project.filePath}</td>
                            <td>{project.gitStats.head.text}</td>
                            <td>{project.gitStats.status.clean}</td>
                        </tr>
                    )
                })}
            </table>
        </div>
    );
}

export default ProjectListing;