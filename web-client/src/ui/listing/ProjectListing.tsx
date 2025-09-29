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
            <table className={"table table-striped"}>
                <tbody>
                {data.projects.map((project, idx) => {
                    return (
                        <tr key={idx}>
                            <td>{project.filePath}</td>
                            <td>{project.gitStats.head.text}</td>
                            <td>{project.gitStats.status.clean}</td>
                            <td><ul>{project.gitStats.remotes.map((remote, remoteIdx) => {
                                return <li key={remoteIdx}>{remote.name}</li>
                            })}</ul></td>
                        </tr>
                    )
                })}
                </tbody>
            </table>
        </div>
    );
}

export default ProjectListing;