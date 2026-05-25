set -e

echo "Staging Changes..."
git add .

read -p "Enter commit message: " message

git commit -m "$message"

read -p "Do you want to push the changes? (y/n): " confirm
if [ "$confirm" = "y" ]; then
    echo "pulling changes from remote..."
    git pull
    echo "pushing changes to remote..."
    git push
else 
    echo "Changes committed but not pushed"
fi