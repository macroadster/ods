package backend

import (
  "fmt"
  "context"
  "github.com/go-git/go-git/v5"
  "github.com/go-git/go-git/v5/plumbing/object"
  "io/ioutil"
  "os"
  "path/filepath"
  "errors"
)

func CreateDAG(ctx context.Context, id string) error {
  // Define the repository directory
  dir := "git_repo"
  err := os.MkdirAll(dir, os.ModePerm)
  if err != nil {
    fmt.Println("Error creating directory:", err)
    return err
  }

  // Open an empty repository
  repo, err := git.PlainInit(dir, false)
  if err != nil {
    fmt.Println("Error creating repository:", err)
    return err
  }
  fn := id+".py"
  dagFile := filepath.Join(".", dir, fn)
  if _, err = os.Stat(dagFile); err == nil {
    fmt.Printf("Error creating dag: %s already exists.\n", dagFile)
    return errors.New("DAG ID already exists.")
  }

  // Create a file with some Python code
  content := []byte("print('Hello from "+id+".py!')")
  err = ioutil.WriteFile(dagFile, content, 0644)
  if err != nil {
    fmt.Println("Error creating file:", err)
    return err
  }

  // Add the file to the staging area (index)
  w, err := repo.Worktree()
  if err != nil {
    fmt.Println("Error accessing worktree:", err)
    return err
  }

  _, err = w.Add(fn)
  if err != nil {
    fmt.Println("Error adding file to index:", err)
    return err
  }

  // Commit the changes with a commit message
  commitMsg := "Added "+id+".py"
  commit, err := w.Commit(commitMsg, &git.CommitOptions{
    Author: &object.Signature{
      Name:  "Your Name",
      Email: "your_email@example.com",
    },
  })
  if err != nil {
    fmt.Println("Error committing changes:", err)
    return err
  }

  fmt.Println("Successfully created repository and committed "+id+".py.")
  fmt.Println("Commit Hash:", commit)
  return nil
}

func DeleteDAG(ctx context.Context, id string) error {
  dir := "git_repo"
  fn := id+".py"
  dagFile := filepath.Join(".", dir, fn)

  // Open existing repository
  repo, err := git.PlainOpen(dir)
  if err != nil {
    fmt.Println("Error opening repository:", err)
    return err
  }

  // Delete the file to the staging area (index)
  w, err := repo.Worktree()
  if err != nil {
    fmt.Println("Error accessing worktree:", err)
    return err
  }

  err = os.Remove(dagFile)
  if err != nil {
    fmt.Println("Error deleting file from file system:", err)
    return err
  }

  // Commit the changes with a commit message
  commitMsg := "Deleted "+id+".py"
  commit, err := w.Commit(commitMsg, &git.CommitOptions{
    Author: &object.Signature{
      Name:  "Your Name",
      Email: "your_email@example.com",
    },
  })
  if err != nil {
    fmt.Println("Error committing changes:", err)
    return err
  }

  fmt.Println("Successfully deleted dag "+id+".py.")
  fmt.Println("Commit Hash:", commit)
  return nil
}

func GetDAG(ctx context.Context, id string) error {
  fmt.Printf("%s\n",id)
  return nil
}
